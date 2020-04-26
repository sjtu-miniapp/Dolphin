## Requirements Specification
**Dolphin** will meet the following requirements:
1. A user can create a group or join one if he's invited.
2. A user can publish a task to a specific group he belongs to. There could be people in that group who are not included in the task.
3. Each task is either a group work or individual work with a deadline.
    + For group work, if a leader is appointed, only he is allowed to change the status of a task to *DONE*. Otherwise, every task worker is equal and can make the change.
    + For individual work, the task workers, everyone would get the same copy of the task, which means every change made to the task is visible to other workers. They have to done their own work and change the *DONE* status.
    + Task can be read only or allowed to be modified by every worker and also the publisher. A read only task can only be modified by the publisher.
4. A calendar is available to every user to see the everyday task and the completion records.
> More to come.

**Dolphin** will **not** meet the following requirements:
1. Publish a group task to several subgroups of people.
2. Share the tasks to other people.