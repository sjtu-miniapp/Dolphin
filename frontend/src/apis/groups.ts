import { cloneDeep } from "lodash";

import { Group } from "../types";

export const SAMPLE_GROUPS: Group[] = [
  { id: "0", name: "商业模式分析", taskNumber: 10, updateTime: new Date() },
  { id: "1", name: "宏观经济学", taskNumber: 1, updateTime: new Date() },
  { id: "2", name: "职业发展规划", taskNumber: 0, updateTime: new Date() }
];

// TODO: remote data call
export async function getGroupsByUser(_userID?: string): Promise<Group[]> {
  const groups = cloneDeep<Group[]>(SAMPLE_GROUPS);
  return new Promise(res => {
    setTimeout(() => res(groups), 20);
  });
}

// TODO: add a group
export async function addGroup(groupName: string): Promise<void> {
  return new Promise(res => {
    setTimeout(() => {
      SAMPLE_GROUPS.push({
        id: SAMPLE_GROUPS.length.toString(),
        name: groupName,
        taskNumber: 0,
        updateTime: new Date()
      });
      res();
    }, 200);
  });
}
