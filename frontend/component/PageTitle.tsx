import Link from 'next/link';
import Image from 'next/image';

export const PageTitle = () => {
  return (
    <Link href="/" className="focus-outline">
      <div className="flex items-center gap-3">
        <Image src="/attendance-plan-icon.png" width={20} height={20} alt="attendance plan icon" className="size-fit" />
        <h1 className="text-xl font-semibold">TOU 受講計画管理（α版）</h1>
      </div>
    </Link>
  );
};
