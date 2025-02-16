import Link from 'next/link';

type Props = {
  href: string;
  children: React.ReactNode;
};

export const LinkText = ({ href, children }: Props) => {
  return (
    <Link href={href} className="text-blue-600 hover:underline">
      {children}
    </Link>
  );
};
