type Props = {
  year: number;
  month: number;
};

export const YearMonthLabel = ({ year, month }: Props) => {
  return (
    <h2 className="text-lg">
      {year}年 {month}月
    </h2>
  );
};
