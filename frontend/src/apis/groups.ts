import { Group } from "src/types";

export const SAMPLE_GROUPS: Group[] = [
  { id: "0", name: "商业模式分析", taskNumber: 10, updateTime: new Date() },
  { id: "1", name: "宏观经济学", taskNumber: 1, updateTime: new Date() },
  { id: "2", name: "职业发展规划", taskNumber: 0, updateTime: new Date() }
];

// TODO: remote data call
export async function getGroupsByUser(_userID?: string): Promise<Group[]> {
  return new Promise(res => {
    setTimeout(() => res(SAMPLE_GROUPS), 20);
  });
}
