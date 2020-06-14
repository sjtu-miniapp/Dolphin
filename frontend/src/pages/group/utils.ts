import * as groupAPI from "../../apis/group";
import * as taskAPI from "../../apis/task";
import { Group } from "../../types";

export const normalizeGroup = async (
  group: groupAPI.GroupShort
): Promise<Group> => {
  const { id, name } = group;

  const tasks = await taskAPI.getTasksByGroupID(id);
  const taskNumber = tasks ? tasks.length : 0;
  const updateTime = new Date();

  return { id, name, taskNumber, updateTime };
};
