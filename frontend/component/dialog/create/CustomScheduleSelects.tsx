'use client';

import { useState } from 'react';

import { SelectContainer } from './SelectContainer';
import { SelectChip } from './SelectChip';
import { AddChip } from './AddChip';

import { useSubject } from '@/provider/SubjectProvider';

type Props = {
  onSelect: (name: string, color: string) => void;
};

export const CustomScheduleSelects = ({ onSelect }: Props) => {
  const [loading, setLoading] = useState(false);
  const { subjects, addSubject, removeSubject, getSubjectByID } = useSubject();

  const click = (id: string) => {
    const subject = getSubjectByID(id);
    if (!subject) return;

    onSelect(subject.name, subject.color);
  };

  const add = async (name: string, color: string) => {
    setLoading(true);
    await addSubject(name, color);
    setLoading(false);
  };

  return (
    <SelectContainer>
      {subjects.map((item, index) => (
        <SelectChip
          key={index}
          label={item.name}
          color={item.color}
          onSelect={() => click(item.id)}
          onRemove={() => removeSubject(item.id)}
        />
      ))}
      <AddChip submit={add} loading={loading} />
    </SelectContainer>
  );
};
