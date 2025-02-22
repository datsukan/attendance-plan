type Props = {
  children: React.ReactNode;
};

export default function EmailLayout({ children }: Props) {
  return (
    <div className="flex min-h-screen items-center justify-center p-2">
      <div className="flex min-h-[25rem] w-[30rem] flex-col items-center gap-8 rounded-lg border p-4 shadow">{children}</div>
    </div>
  );
}
