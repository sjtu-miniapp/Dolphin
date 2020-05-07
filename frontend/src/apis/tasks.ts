import { Task } from "src/types";

export const SAMPLE_TASKS: Task[] = [
  {
    id: "0",
    groupID: "0",
    name: "Lab 1",
    description: `作业一
小组作业
自由组队
2~3人
2020-06-01 00:00 前上传课题研究报告到ftp
ftp域名: ftp.antai.sjtu.com/business-pattern/01
用户名: student
密码: antai-rw-2020
课题: 华为商业模式研究
`,
    publisher: "TA",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,
    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Bob1", status: "completed" },
      { userName: "Bob2", status: "completed" },
      { userName: "Bob3", status: "completed" },
      { userName: "Bobo", status: "completed" },
      { userName: "Bob1", status: "completed" },
      { userName: "Bob2", status: "completed" },
      { userName: "Bob3", status: "completed" },
      { userName: "Bobo", status: "completed" },
      { userName: "Bob1", status: "completed" },
      { userName: "Bob2", status: "completed" },
      { userName: "Bob3", status: "completed" },
      { userName: "Bobo", status: "completed" },
      { userName: "Bob1", status: "completed" },
      { userName: "Bob2", status: "completed" },
      { userName: "Bob3", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "Complete"
  },
  {
    id: "1",
    groupID: "0",
    name: "Lab 2",
    description: "Make Progress",
    publisher: "TA",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "In Progess"
  },
  {
    id: "2",
    groupID: "0",
    name: "Lab 3",
    description: "Review",
    publisher: "Pro",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "Todo"
  },
  {
    id: "3",
    groupID: "1",
    name: "Project 1",
    description: "Warm up",
    publisher: "TA",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "Complete"
  },
  {
    id: "4",
    groupID: "1",
    name: "Project 2",
    description: "Make Progress",
    publisher: "TA",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "In Progess"
  },
  {
    id: "5",
    groupID: "1",
    name: "Project 3",
    description: "Review",
    publisher: "Pro",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "Todo"
  },
  {
    id: "6",
    groupID: "2",
    name: "Paper 1",
    description: "Warm up",
    publisher: "TA",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "Complete"
  },
  {
    id: "7",
    groupID: "2",
    name: "Paper 2",
    description: "Make Progress",
    publisher: "TA",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,

    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "In Progess"
  },
  {
    id: "8",
    groupID: "2",
    name: "Paper 3",
    description: "Review",
    publisher: "Pro",
    startDate: new Date(),
    endDate: new Date(),
    readOnly: true,
    receivers: [
      { userName: "Bobo", status: "completed" },
      { userName: "Biubiu", status: "in-progress" },
      { userName: "Piupiu", status: "to-start" }
    ],
    status: "Todo"
  }
];

// TODO: remote data call
export async function getTasksByGroup(groupID: string): Promise<Task[]> {
  return new Promise(res => {
    setTimeout(() => res(SAMPLE_TASKS.filter(t => t.groupID === groupID)), 20);
  });
}

export async function getTaskByID(id: string): Promise<Task | undefined> {
  return new Promise(res => {
    setTimeout(() => res(SAMPLE_TASKS.find(t => t.id === id)), 20);
  });
}
