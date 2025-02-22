type Props = {
  label: string;
};

export const FormTitle = ({ label }: Props) => {
  return <h2 className="text-2xl">{label}</h2>;
};
