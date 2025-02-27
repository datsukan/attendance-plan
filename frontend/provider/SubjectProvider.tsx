'use client';

import { createContext, useContext, useState, useEffect } from 'react';
import toast from 'react-hot-toast';

import type { Type } from '@/type';
import { fetchSubjects } from '@/backend-api/fetchSubjects';
import { createSubject } from '@/backend-api/createSubject';
import { deleteSubject } from '@/backend-api/deleteSubject';
import { loadAuthUser } from '@/storage/user';

type SubjectContextType = {
  subjects: Type.Subject[];
  addSubject: (name: string, color: string) => Promise<void>;
  removeSubject: (id: string) => Promise<void>;
  getSubjectByID: (id: string) => Type.Subject | undefined;
};

const createCtx = () => {
  const ctx = createContext<SubjectContextType | undefined>(undefined);
  const useCtx = () => {
    const c = useContext(ctx);
    if (!c) throw new Error('useCtx must be inside a Provider with a value');
    return c;
  };
  return [useCtx, ctx.Provider] as const;
};

const [useCtx, SetSubjectProvider] = createCtx();
export const useSubject = useCtx;

type Props = {
  children: React.ReactNode;
};

export const SubjectProvider = ({ children }: Props) => {
  const [subjects, setSubjects] = useState<Type.Subject[]>([]);

  useEffect(() => {
    const au = loadAuthUser();
    if (!au) return;

    (async () => {
      try {
        const subjects = await fetchSubjects();
        setSubjects(subjects);
      } catch (e) {
        toast.error('科目情報の取得に失敗しました');
        toast.error(String(e));
        return;
      }
    })();
  }, []);

  const addSubject = async (name: string, color: string) => {
    try {
      const newSubject = await createSubject(name, color);
      setSubjects([...subjects, newSubject]);
    } catch (e) {
      toast.error('科目の追加に失敗しました');
      toast.error(String(e));
      return;
    }
  };

  const removeSubject = async (id: string) => {
    try {
      await deleteSubject(id);
      setSubjects(subjects.filter((s) => s.id !== id));
    } catch (e) {
      toast.error('科目の削除に失敗しました');
      toast.error(String(e));
      return;
    }
  };

  const getSubjectByID = (id: string) => subjects.find((s) => s.id === id);

  return <SetSubjectProvider value={{ subjects, addSubject, removeSubject, getSubjectByID }}>{children}</SetSubjectProvider>;
};
