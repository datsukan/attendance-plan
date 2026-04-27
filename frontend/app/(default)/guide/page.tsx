import type { Metadata } from 'next';
import Link from 'next/link';
import Image from 'next/image';
import {
  CalendarDaysIcon,
  PlusCircleIcon,
  CursorArrowRaysIcon,
  Squares2X2Icon,
  ArrowUturnLeftIcon,
  SwatchIcon,
  DocumentDuplicateIcon,
  DevicePhoneMobileIcon,
} from '@heroicons/react/24/outline';

export const metadata: Metadata = {
  title: 'このツールについて',
  description: 'TOU 受講スケジュール管理の機能・使い方についての説明ページです。',
};

const features = [
  {
    icon: CalendarDaysIcon,
    title: 'カレンダー表示',
    description: '月単位のカレンダーで受講スケジュールを一目で確認できます。無限スクロールで過去・未来の月をスムーズに閲覧できます。',
  },
  {
    icon: PlusCircleIcon,
    title: 'スケジュール作成',
    description: '「学事」と「受講」の2種類のスケジュールを作成できます。名前・日付・期間・色を自由に設定できます。',
  },
  {
    icon: CursorArrowRaysIcon,
    title: 'ドラッグ&ドロップ',
    description: 'スケジュールをカレンダー上でドラッグするだけで日程を変更できます。直感的な操作でスケジュール調整が簡単です。',
  },
  {
    icon: Squares2X2Icon,
    title: '一括操作',
    description: '複数のスケジュールを同時に選択して、まとめて削除・移動が可能です。大量のスケジュール変更も素早く処理できます。',
  },
  {
    icon: ArrowUturnLeftIcon,
    title: '取り消し機能',
    description: '操作を間違えても取り消しボタンで直前の状態に戻せます。誤って削除・移動してしまっても安心して作業できます。',
  },
  {
    icon: SwatchIcon,
    title: 'カラー設定',
    description: 'スケジュールに色を設定して視覚的に管理できます。種別や優先度をカラーで区別することで、カレンダーが一層見やすくなります。',
  },
  {
    icon: DocumentDuplicateIcon,
    title: 'まとめて作成',
    description: '繰り返しのスケジュールを一括登録できます。テンプレートを使った一括作成で、定期的な受講スケジュールも素早く登録できます。',
  },
  {
    icon: DevicePhoneMobileIcon,
    title: 'PC・モバイル対応',
    description: 'パソコン・スマートフォン・タブレットのどのデバイスからでも利用できます。外出先でもスケジュールを確認・管理できます。',
  },
];

const steps = [
  {
    number: '01',
    title: 'アカウントを作成する',
    description: 'メールアドレスとパスワードでアカウントを作成します。すでにアカウントをお持ちの場合はサインインしてください。',
  },
  {
    number: '02',
    title: 'スケジュールを追加する',
    description: 'カレンダー上部の「作成」ボタンをクリックして、種類（学事・受講）・名前・日付・期間・色を設定して登録します。繰り返しスケジュールはまとめて登録することもできます。',
  },
  {
    number: '03',
    title: 'カレンダーで確認・管理する',
    description: 'カレンダーに登録したスケジュールが表示されます。スケジュールをクリックすると詳細の確認や編集・削除ができます。',
  },
  {
    number: '04',
    title: 'ドラッグ&ドロップで調整する',
    description: 'スケジュールを別の日付にドラッグするだけで日程を変更できます。間違えたときは「取り消し」ボタンですぐに元に戻せます。',
  },
];

export default function GuidePage() {
  return (
    <div className="mx-auto max-w-4xl space-y-16 py-8">
      {/* Hero Section */}
      <section className="flex flex-col items-center gap-6 text-center">
        <div className="flex items-center gap-3">
          <Image src="/attendance-plan-icon.png" width={48} height={48} alt="attendance plan icon" className="size-fit" />
          <h2 className="text-3xl font-bold">TOU 受講スケジュール管理</h2>
        </div>
        <p className="max-w-xl text-lg text-gray-600">
          直感的で手軽に受講スケジュールを管理するツール
        </p>
        <p className="max-w-2xl text-gray-500">
          TOU の受講スケジュールをカレンダー形式で視覚的に管理できます。
          ドラッグ&ドロップや一括操作など、スケジュール管理を快適にする機能を揃えています。
        </p>
        <Link
          href="/"
          className="rounded-full bg-blue-600 px-8 py-3 font-medium text-white hover:bg-blue-500 active:bg-blue-400"
        >
          カレンダーを開く
        </Link>
      </section>

      {/* About Section */}
      <section className="rounded-xl border bg-gray-50 p-8">
        <h3 className="mb-4 text-xl font-semibold">このツールについて</h3>
        <div className="space-y-3 text-gray-600">
          <p>
            TOU 受講スケジュール管理は、東京通信大学（TOU）の受講スケジュールを
            カレンダー形式で直感的に管理するためのツールです。
          </p>
          <p>
            受講する授業・講義の日程をカレンダーに登録することで、全体のスケジュールを一目で把握できます。
            日程の変更や調整もドラッグ&ドロップで簡単に行えます。
          </p>
          <p>
            大学が定める授業配信・試験などの「学事」スケジュールと、自分が受講する授業の「受講」スケジュールの
            2 種類を使い分けることで、複雑なスケジュールもわかりやすく整理できます。
          </p>
        </div>
      </section>

      {/* Usage Example */}
      <div className="overflow-hidden rounded-xl border shadow-sm">
        <p className="border-b bg-gray-50 px-5 py-3 text-sm font-medium text-gray-600">使用例</p>
        <Image
          src="/usage-example.png"
          width={1200}
          height={800}
          alt="カレンダーの使用例"
          className="w-full"
        />
      </div>

      {/* Features Section */}
      <section>
        <h3 className="mb-8 text-center text-xl font-semibold">主な機能</h3>
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {features.map((feature) => (
            <div
              key={feature.title}
              className="flex flex-col gap-3 rounded-xl border bg-white p-5 shadow-sm"
            >
              <feature.icon className="size-7 text-blue-600" />
              <h4 className="font-semibold">{feature.title}</h4>
              <p className="text-sm text-gray-500">{feature.description}</p>
            </div>
          ))}
        </div>
      </section>

      {/* Two-track Section */}
      <section>
        <h3 className="mb-3 text-center text-xl font-semibold">学事と受講を分けて表示</h3>
        <p className="mb-8 text-center text-gray-500">
          各日付には「学事」エリアと「受講」エリアが固定分離されており、2 種類のスケジュールを混在なく整理して確認できます。
        </p>
        <div className="overflow-hidden rounded-xl border bg-white shadow-sm">
          <div className="flex flex-col items-stretch divide-y sm:flex-row sm:divide-x sm:divide-y-0">
            <div className="flex items-center justify-center bg-gray-50 p-8 sm:w-72">
              <div className="w-56 overflow-hidden rounded-xl border-2 border-gray-300 text-sm shadow">
                <div className="border-b bg-white px-3 py-1.5 text-right font-semibold text-gray-600">15</div>
                <div className="space-y-1.5 border-b bg-yellow-50 p-2">
                  <div className="rounded bg-red-500 px-2 py-0.5 text-xs font-medium text-white">履修登録</div>
                  <div className="rounded bg-yellow-400 px-2 py-0.5 text-xs font-medium text-white">第3回 授業配信</div>
                </div>
                <div className="space-y-1.5 bg-white p-2">
                  <div className="rounded border-2 border-orange-300 px-2 py-0.5 text-xs font-medium text-orange-600">第1回 数学基礎</div>
                  <div className="rounded bg-blue-500 px-2 py-0.5 text-xs font-medium text-white">第2回 経営学演習</div>
                </div>
              </div>
            </div>
            <div className="flex flex-col justify-center gap-6 p-8">
              <div className="flex items-start gap-4">
                <div className="flex h-8 w-10 shrink-0 items-center justify-center self-center rounded-full bg-yellow-100 text-xs font-bold text-yellow-700">学事</div>
                <div>
                  <p className="font-semibold text-gray-800">上段：学事スケジュール</p>
                  <p className="mt-1 text-sm text-gray-500">授業配信日・試験日・成績発表日など、大学が定めるスケジュールを上段に固定表示します。</p>
                </div>
              </div>
              <div className="flex items-start gap-4">
                <div className="flex h-8 w-10 shrink-0 items-center justify-center self-center rounded-full bg-blue-100 text-xs font-bold text-blue-700">受講</div>
                <div>
                  <p className="font-semibold text-gray-800">下段：受講スケジュール</p>
                  <p className="mt-1 text-sm text-gray-500">自分が立てた受講計画を下段に固定表示します。学事との対比を見ながら、効率よく計画を組み立てられます。</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* How to Use Section */}
      <section>
        <h3 className="mb-8 text-center text-xl font-semibold">使い方</h3>
        <div className="space-y-4">
          {steps.map((step, index) => (
            <div key={step.number} className="flex gap-5 rounded-xl border bg-white p-6 shadow-sm">
              <div className="flex size-10 shrink-0 items-center justify-center self-center rounded-full bg-blue-600 text-sm font-bold text-white">
                {index + 1}
              </div>
              <div className="space-y-1">
                <h4 className="font-semibold">{step.title}</h4>
                <p className="text-sm text-gray-500">{step.description}</p>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Bottom CTA */}
      <section className="flex flex-col items-center gap-4 rounded-xl border bg-blue-50 p-10 text-center">
        <h3 className="text-xl font-semibold">さっそく使ってみる</h3>
        <p className="text-gray-600">カレンダーページから受講スケジュールの管理をはじめましょう。</p>
        <Link
          href="/"
          className="rounded-full bg-blue-600 px-8 py-3 font-medium text-white hover:bg-blue-500 active:bg-blue-400"
        >
          カレンダーを開く
        </Link>
      </section>
    </div>
  );
}
