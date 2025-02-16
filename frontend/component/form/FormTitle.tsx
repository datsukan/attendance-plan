type Props = {
  label: string;
};

export const FormTitle = ({ label }: Props) => {
  return <h1 className="text-2xl">{label}</h1>;
};
