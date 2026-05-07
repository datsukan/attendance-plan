'use client';

import { useEffect, useState } from 'react';
import { ChevronDownIcon, ChevronRightIcon } from '@heroicons/react/20/solid';

import { useAuth } from '@/hook/useAuth';
import { useUser } from '@/provider/UserProvider';
import { fetchUserUsages, UserUsage } from '@/backend-api/fetchUserUsages';

export default function Usage() {
  const [loadedAuth, isAuth] = useAuth();
  const { user } = useUser();
  const [usages, setUsages] = useState<UserUsage[]>([]);
  const [expandedIds, setExpandedIds] = useState<Set<string>>(new Set());

  const adminEmails = (process.env.NEXT_PUBLIC_ADMIN_EMAILS ?? '').split(',').map((e) => e.trim()).filter(Boolean);
  const isAdmin = user ? adminEmails.includes(user.email) : false;

  useEffect(() => {
    if (!isAuth || !isAdmin) return;

    (async () => {
      try {
        const data = await fetchUserUsages();
        setUsages(data);
      } catch {
        // エラーは error.ts のグローバルハンドラに委ねる
      }
    })();
  }, [isAuth, isAdmin]);

  if (!loadedAuth || !isAuth) return null;

  if (!isAdmin) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <p className="text-gray-500">アクセス権限がありません</p>
      </div>
    );
  }

  const toggleExpand = (id: string) => {
    setExpandedIds((prev) => {
      const next = new Set(prev);
      if (next.has(id)) {
        next.delete(id);
      } else {
        next.add(id);
      }
      return next;
    });
  };

  return (
    <div className="p-6">
      <h2 className="mb-4 text-2xl font-semibold">利用状況</h2>
      <p className="mb-6 text-sm text-gray-500">登録ユーザー数: {usages.length}</p>

      <div className="overflow-x-auto rounded-lg border shadow-sm">
        <table className="min-w-full divide-y divide-gray-200 text-sm">
          <thead className="bg-gray-50">
            <tr>
              <th className="w-8 px-4 py-3" />
              <th className="px-4 py-3 text-left font-medium text-gray-600">名前</th>
              <th className="px-4 py-3 text-left font-medium text-gray-600">メールアドレス</th>
              <th className="px-4 py-3 text-left font-medium text-gray-600">登録日時</th>
              <th className="px-4 py-3 text-left font-medium text-gray-600">最終利用日時</th>
              <th className="px-4 py-3 text-right font-medium text-gray-600">テンプレート科目数</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-100 bg-white">
            {usages.map((u) => (
              <>
                <tr key={u.id} className="cursor-pointer hover:bg-gray-50" onClick={() => toggleExpand(u.id)}>
                  <td className="px-4 py-3 text-gray-400">
                    {expandedIds.has(u.id) ? (
                      <ChevronDownIcon className="size-4" />
                    ) : (
                      <ChevronRightIcon className="size-4" />
                    )}
                  </td>
                  <td className="px-4 py-3">{u.name || '—'}</td>
                  <td className="px-4 py-3 text-gray-600">{u.email}</td>
                  <td className="px-4 py-3 text-gray-600">{u.registered_at}</td>
                  <td className="px-4 py-3 text-gray-600">{u.last_used_at}</td>
                  <td className="px-4 py-3 text-right">{u.subjects.length}</td>
                </tr>
                {expandedIds.has(u.id) && u.subjects.length > 0 && (
                  <tr key={`${u.id}-subjects`}>
                    <td colSpan={6} className="bg-gray-50 px-8 py-3">
                      <p className="mb-2 text-xs font-medium text-gray-500">テンプレート科目</p>
                      <table className="min-w-full text-xs">
                        <thead>
                          <tr className="text-gray-500">
                            <th className="pb-1 text-left font-medium">科目名</th>
                            <th className="pb-1 text-left font-medium">カラー</th>
                            <th className="pb-1 text-left font-medium">最終更新日時</th>
                          </tr>
                        </thead>
                        <tbody className="divide-y divide-gray-200">
                          {u.subjects.map((s) => (
                            <tr key={s.id}>
                              <td className="py-1 pr-6">{s.name}</td>
                              <td className="py-1 pr-6">
                                <span className="inline-flex items-center gap-1">
                                  <span
                                    className="inline-block size-3 rounded-full border"
                                    style={{ backgroundColor: s.color }}
                                  />
                                  {s.color}
                                </span>
                              </td>
                              <td className="py-1 text-gray-500">{s.updated_at}</td>
                            </tr>
                          ))}
                        </tbody>
                      </table>
                    </td>
                  </tr>
                )}
                {expandedIds.has(u.id) && u.subjects.length === 0 && (
                  <tr key={`${u.id}-empty`}>
                    <td colSpan={6} className="bg-gray-50 px-8 py-3 text-xs text-gray-400">
                      テンプレート科目なし
                    </td>
                  </tr>
                )}
              </>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
