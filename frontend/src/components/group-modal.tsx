import Taro, { FC, useState } from '@tarojs/taro';
import { AtModal, AtModalHeader, AtModalAction, AtModalContent, AtInput } from 'taro-ui';
import { Button } from '@tarojs/components';

interface GroupModalProps {
  isOpened: boolean;
  handleClose: () => void;
  handleCancel: () => void;
  handleConfirm: (groupName: string) => void;
}

const GroupModal: FC<GroupModalProps> = props => {
  const [inputValue, setInputValue] = useState<string>('');

  const updateInputValue = (v: string, _e: any) => {
    setInputValue(v);
  }

  const onTriggerConfirm = () => {
    if (!inputValue) return;
    props.handleConfirm(inputValue);
  }

  return (
    <AtModal isOpened={props.isOpened} onClose={props.handleClose}>
      <AtModalHeader>创建小组</AtModalHeader>
      <AtModalContent>
        <AtInput name='groupName' title='组名' type='text' maxLength={20} placeholder='请输入小组名' value={inputValue} onChange={updateInputValue} />
      </AtModalContent>
      <AtModalAction> <Button onClick={props.handleCancel}>取消</Button> <Button onClick={onTriggerConfirm}>确定</Button> </AtModalAction>
    </AtModal>

  )
}

GroupModal.defaultProps = {
  isOpened: false,
  handleClose: () => { },
  handleCancel: () => { },
  handleConfirm: (_groupName: string) => { }
}

export default GroupModal;
