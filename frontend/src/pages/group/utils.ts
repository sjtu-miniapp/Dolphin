import * as groupAPI from "../../apis/group";
import * as taskAPI from "../../apis/task";
import { Group } from "../../types";

export const normalizeGroup = async (
  group: groupAPI.GroupShort
): Promise<Group> => {
  const { id, name } = group;

  const tasks = await taskAPI.getTasksByGroupID(id);
  const taskNumber = tasks ? tasks.length : 0;

  const [d, t] = (group.updated_at || "").split(" ");
  console.log(group.updated_at, d, t);

  const updateTime = new Date(`${d}T${t}.000Z`);
  console.log("uuuuuuuuuuu", updateTime);
  return { id, name, taskNumber, updateTime };
};
