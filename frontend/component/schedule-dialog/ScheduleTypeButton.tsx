type Props = {
  label: string;
  isSelected: boolean;
  onClick: () => void;
};

export const ScheduleTypeButton = ({ label, isSelected, onClick }: Props) => {
  const dynamicClassName = isSelected
    ? 'bg-blue-100 hover:bg-blue-200 active:bg-blue-300 text-blue-600'
    : 'hover:bg-gray-100 active:bg-gray-200 text-gray-600';
  return (
    <button className={`px-3 py-1 rounded ${dynamicClassName}`} onClick={onClick}>
      {label}
    </button>
  );
};
