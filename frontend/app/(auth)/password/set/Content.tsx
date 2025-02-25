'use client';

import { useState, Suspense } from 'react';

import { Form } from './Form';
import { Note } from './Note';
import { CompleteMessage } from './CompleteMessage';

export const Content = () => {
  const [isComplete, setIsComplete] = useState(false);

  if (isComplete) {
    return (
      <div className="pt-16">
        <CompleteMessage />
      </div>
    );
  } else {
    return (
      <div className="flex flex-col gap-8">
        <Suspense>
          <Form complete={() => setIsComplete(true)} />
        </Suspense>
        <Note />
      </div>
    );
  }
};
