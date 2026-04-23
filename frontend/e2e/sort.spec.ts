import { test, expect, type Page } from '@playwright/test';

// ─── 定数 ────────────────────────────────────────────────────────────────────

const FAKE_USER = {
  id: 'test-user-id',
  email: 'test@example.com',
  name: 'テストユーザー',
  session_token: 'fake-session-token',
};

// ─── ヘルパー ─────────────────────────────────────────────────────────────────

/** localStorage にダミー認証情報を注入して useAuth() のリダイレクトをバイパスする */
async function setupAuth(page: Page) {
  await page.addInitScript((user) => {
    window.localStorage.setItem('auth-user', JSON.stringify(user));
  }, FAKE_USER);
}

/**
 * バックエンド API をモックする。
 * - GET  .../users/:id           → ダミーユーザーを返す
 * - GET  .../users/:id/schedules → mockBody を返す
 * - GET  .../users/:id/subjects  → 空リストを返す
 * - PUT  .../schedules/bulk      → 200 を返す（ソート保存を受け入れる）
 */
async function mockApi(page: Page, scheduleBody: object) {
  await page.route('**/users/*/schedules', async (route) => {
    await route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(scheduleBody),
    });
  });

  await page.route('**/users/*/subjects', async (route) => {
    await route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ subjects: [] }),
    });
  });

  await page.route('**/schedules/bulk', async (route) => {
    await route.fulfill({ status: 200, contentType: 'application/json', body: '{}' });
  });

  // ユーザー取得 API（getUser）
  await page.route(`**/users/${FAKE_USER.id}`, async (route) => {
    if (route.request().method() === 'GET') {
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          id: FAKE_USER.id,
          email: FAKE_USER.email,
          name: FAKE_USER.name,
          created_at: '2026-01-01T00:00:00.000Z',
          updated_at: '2026-01-01T00:00:00.000Z',
        }),
      });
    } else {
      await route.fallback();
    }
  });
}

/**
 * dnd-kit の PointerSensor（activationConstraint: distance:5）に対応したドラッグ。
 * Playwright の mouse API で pointer events を直接制御する。
 * 要素をビューポート内にスクロールしてから activationConstraint: distance:5 を
 * 超える移動を行い、確実にドラッグを有効化する。
 */
async function dragTo(page: Page, sourceLocator: ReturnType<Page['locator']>, targetLocator: ReturnType<Page['locator']>) {
  // ドラッグ前に要素をビューポート内にスクロールする
  await sourceLocator.scrollIntoViewIfNeeded();
  await targetLocator.scrollIntoViewIfNeeded();

  const sourceBox = await sourceLocator.boundingBox();
  const targetBox = await targetLocator.boundingBox();
  if (!sourceBox || !targetBox) throw new Error('ドラッグ対象の要素が見つかりません');

  const sx = sourceBox.x + sourceBox.width / 2;
  const sy = sourceBox.y + sourceBox.height / 2;
  const tx = targetBox.x + targetBox.width / 2;
  const ty = targetBox.y + targetBox.height / 2;

  // PointerSensor の activationConstraint: distance:5 を超えるために
  // 垂直方向（同列内ソートなので Y 軸方向）に十分な距離を移動してから
  // ターゲットへ向かう。各ステップを細かくして pointermove イベントを確実に発火させる。
  await page.mouse.move(sx, sy);
  await page.mouse.down();
  await page.mouse.move(sx, sy + 1, { steps: 1 });
  await page.mouse.move(sx, sy + 3, { steps: 2 });
  await page.mouse.move(sx, sy + 6, { steps: 3 });
  await page.mouse.move(sx, sy + 10, { steps: 4 });
  await page.mouse.move(tx, ty, { steps: 20 });
  await page.waitForTimeout(200);
  await page.mouse.up();
  await page.waitForTimeout(1000);
}

// ─── テストデータ ─────────────────────────────────────────────────────────────

// 今日の週（2026-04-20 の週）に単一日付スケジュールを2件配置
const singleDaySchedules = {
  master_schedules: [
    {
      date: '2026-04-22',
      type: 'master',
      schedules: [
        {
          id: 'single-1',
          name: '予定A',
          starts_at: '2026-04-22T00:00:00.000Z',
          ends_at: '2026-04-22T00:00:00.000Z',
          color: 'blue',
          type: 'master',
          order: 1,
        },
        {
          id: 'single-2',
          name: '予定B',
          starts_at: '2026-04-22T00:00:00.000Z',
          ends_at: '2026-04-22T00:00:00.000Z',
          color: 'red',
          type: 'master',
          order: 2,
        },
      ],
    },
  ],
  custom_schedules: [],
};

// 同じ開始日から複数日にまたがるスケジュールを2件配置
const rangeDateSchedules = {
  master_schedules: [
    {
      date: '2026-04-21',
      type: 'master',
      schedules: [
        {
          id: 'range-1',
          name: '複数日A',
          starts_at: '2026-04-21T00:00:00.000Z',
          ends_at: '2026-04-23T00:00:00.000Z',
          color: 'blue',
          type: 'master',
          order: 1,
        },
        {
          id: 'range-2',
          name: '複数日B',
          starts_at: '2026-04-21T00:00:00.000Z',
          ends_at: '2026-04-23T00:00:00.000Z',
          color: 'red',
          type: 'master',
          order: 2,
        },
      ],
    },
  ],
  custom_schedules: [],
};

// 週跨ぎスケジュール（日→月）と同週の別スケジュール
const crossWeekSchedules = {
  master_schedules: [
    {
      date: '2026-04-19',
      type: 'master',
      schedules: [
        {
          id: 'cross-1',
          name: '跨週スケジュール',
          starts_at: '2026-04-19T00:00:00.000Z',
          ends_at: '2026-04-21T00:00:00.000Z',
          color: 'blue',
          type: 'master',
          order: 1,
        },
      ],
    },
    {
      date: '2026-04-20',
      type: 'master',
      schedules: [
        {
          id: 'cross-2',
          name: '単日スケジュール',
          starts_at: '2026-04-20T00:00:00.000Z',
          ends_at: '2026-04-20T00:00:00.000Z',
          color: 'red',
          type: 'master',
          order: 1,
        },
      ],
    },
  ],
  custom_schedules: [],
};

// ─── テスト ───────────────────────────────────────────────────────────────────

test.describe('ソート: 単一日付スケジュール', () => {
  test.beforeEach(async ({ page }) => {
    await setupAuth(page);
    await mockApi(page, singleDaySchedules);
    await page.goto('/');
    await page.waitForSelector('text=予定A', { timeout: 10000 });
    await page.waitForSelector('text=予定B', { timeout: 10000 });
  });

  test('2件のスケジュールが表示される', async ({ page }) => {
    await expect(page.getByText('予定A').first()).toBeVisible();
    await expect(page.getByText('予定B').first()).toBeVisible();
  });

  test('予定Aを予定Bの上にドラッグするとAPIが呼ばれる', async ({ page }) => {
    const allRequests: string[] = [];
    const consoleMsgs: string[] = [];
    const errors: string[] = [];
    page.on('request', (req) => allRequests.push(`${req.method()} ${req.url()}`));
    page.on('console', (msg) => consoleMsgs.push(`[${msg.type()}] ${msg.text()}`));
    page.on('pageerror', (e) => errors.push(e.message));

    const putRequest = page.waitForRequest(
      (req) => req.url().includes('/schedules/bulk') && req.method() === 'PUT',
      { timeout: 10000 }
    );

    await dragTo(
      page,
      page.getByText('予定A').first(),
      page.getByText('予定B').first()
    );

    try {
      const request = await putRequest;
      expect(request).toBeTruthy();
    } catch (e) {
      console.log('=== 全リクエスト ===\n', allRequests.join('\n'));
      console.log('=== コンソール ===\n', consoleMsgs.join('\n'));
      console.log('=== エラー ===\n', errors.join('\n'));
      throw e;
    }
  });

});

test.describe('ソート: 範囲日付スケジュール', () => {
  test.beforeEach(async ({ page }) => {
    await setupAuth(page);
    await mockApi(page, rangeDateSchedules);
    await page.goto('/');
    await page.waitForSelector('text=複数日A', { timeout: 10000 });
    await page.waitForSelector('text=複数日B', { timeout: 10000 });
  });

  test('2件の範囲日付スケジュールが表示される', async ({ page }) => {
    await expect(page.getByText('複数日A').first()).toBeVisible();
    await expect(page.getByText('複数日B').first()).toBeVisible();
  });

  test('複数日Aを複数日Bの上にドラッグするとAPIが呼ばれる', async ({ page }) => {
    const putRequest = page.waitForRequest(
      (req) => req.url().includes('/schedules/bulk') && req.method() === 'PUT',
      { timeout: 10000 }
    );

    await dragTo(
      page,
      page.getByText('複数日A').first(),
      page.getByText('複数日B').first()
    );

    const request = await putRequest;
    expect(request).toBeTruthy();
  });

});

test.describe('ソート: 週跨ぎスケジュール', () => {
  test.beforeEach(async ({ page }) => {
    await setupAuth(page);
    await mockApi(page, crossWeekSchedules);
    await page.goto('/');
    await page.waitForSelector('text=跨週スケジュール', { timeout: 10000 });
  });

  test('週跨ぎスケジュールが表示される', async ({ page }) => {
    await expect(page.getByText('跨週スケジュール').first()).toBeVisible();
  });

  test('週跨ぎスケジュールを操作しても JS エラーが発生しない', async ({ page }) => {
    const errors: string[] = [];
    page.on('pageerror', (e) => errors.push(e.message));

    // 跨週スケジュールをクリック（ドラッグではなく操作できることを確認）
    await page.getByText('跨週スケジュール').first().click();

    // dnd-kit 関連のエラーが発生していないことを確認
    const dndErrors = errors.filter((e) =>
      e.toLowerCase().includes('sortable') ||
      e.toLowerCase().includes('dnd') ||
      e.toLowerCase().includes('duplicate')
    );
    expect(dndErrors).toHaveLength(0);
  });
});
