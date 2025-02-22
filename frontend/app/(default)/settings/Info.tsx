import { Email } from './Email';
import { Name } from './Name';
import { DeleteButton } from './DeleteButton';

export const Info = () => {
  return (
    <div className="flex w-full flex-col justify-between gap-8">
      <div className="flex w-full flex-col gap-8">
        <Email />
        <Name />
      </div>
      <DeleteButton />
    </div>
  );
};
