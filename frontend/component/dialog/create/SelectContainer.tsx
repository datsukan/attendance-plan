type Props = {
  children: React.ReactNode;
};

export const SelectContainer = ({ children }: Props) => {
  return <div className="flex flex-wrap gap-1">{children}</div>;
};
