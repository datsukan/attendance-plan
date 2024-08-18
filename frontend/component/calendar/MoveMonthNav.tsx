import { NavButton } from './NavButton';
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/react/24/outline';

type Props = {
  prev: () => void;
  next: () => void;
  reset: () => void;
};

export const MoveMonthNav = ({ prev, next, reset }: Props) => {
  return (
    <div className="flex gap-2">
      <NavButton onClick={() => prev()}>
        <ChevronLeftIcon className="size-5" />
      </NavButton>
      <NavButton onClick={() => reset()}>
        <div className="flex justify-center">
          <span className="text-sm">今日</span>
        </div>
      </NavButton>
      <NavButton onClick={() => next()}>
        <ChevronRightIcon className="size-5" />
      </NavButton>
    </div>
  );
};
