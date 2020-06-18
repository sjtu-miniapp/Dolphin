import Taro, { FC, useState, useDidShow } from '@tarojs/taro';
import { AtModal, AtModalAction, AtModalContent, AtInput, AtSegmentedControl } from 'taro-ui';
import { Button } from '@tarojs/components';

export type InputType = 'groupName' | 'shareCode';

interface GroupModalProps {
  isOpened: boolean;
  handleClose: () => void;
  handleCancel: () => void;
  handleConfirm: (inputType: InputType, value: string) => void;
}

const GroupModal: FC<GroupModalProps> = props => {
  const [inputType, setInputType] = useState<InputType>('groupName');

  const [inputValue, setInputValue] = useState<string>('');

  useDidShow(() => setInputValue(''));

  const updateInputValue = (v: string, _e: any) => {
    setInputValue(v);
  }

  const onTriggerConfirm = () => {
    if (!inputValue) return;
    const value = inputValue;
    setInputValue('');
    props.handleConfirm(inputType, value);
  }

  const title = inputType === 'groupName' ? '组名' : '邀请码';

  const switchInputType = v => {
    if (v === 0) setInputType('groupName');
    else setInputType('shareCode');
  }

  return (
    <AtModal isOpened={props.isOpened} onClose={props.handleClose}>
      <AtModalContent>
        <AtSegmentedControl
          values={['创建小组', '加入小组']}
          onClick={switchInputType}
          current={inputType === 'groupName' ? 0 : 1}
        />
        <AtInput
          name='toInput'
          title={title}
          type='text'
          maxLength={20}
          value={inputValue}
          onChange={updateInputValue}
          required
          focus={props.isOpened}
        />
      </AtModalContent>
      <AtModalAction> <Button onClick={props.handleCancel}>取消</Button> <Button onClick={onTriggerConfirm}>确定</Button> </AtModalAction>
    </AtModal>

  )
}

GroupModal.defaultProps = {
  isOpened: false,
  handleClose: () => { },
  handleCancel: () => { },
  handleConfirm: (_type: 'groupName', _groupName: string) => { }
}

export default GroupModal;
