import { SelectContainer } from './SelectContainer';
import { SelectChip } from './SelectChip';
import { getMasterScheduleTemplates } from '@/component/schedule/schedule-module';

type Props = {
  onSelect: (name: string, color: string) => void;
};

export const MasterScheduleTemplates = ({ onSelect }: Props) => {
  const templates = getMasterScheduleTemplates();

  return (
    <SelectContainer>
      {templates.map((item, index) => (
        <SelectChip key={index} label={item.name} color={item.color} onSelect={() => onSelect(item.name, item.color)} />
      ))}
    </SelectContainer>
  );
};
