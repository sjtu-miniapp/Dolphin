## Requirements Specification

**Dolphin** will meet the following requirements:

1. A user can create a group or join one if he's invited.
2. A user can publish a task to a specific group he belongs to. There could be people in that group who are not included in the task.
3. Each task is either a group work or individual work with a deadline.
   - For group work, every one belong to the group are available to change the status to _DONE_, . Each operation should be recorded so that everyone in the group would know the historical changes. (Like how we tracking by git.)
   - For individual work, the task workers, everyone would get the same copy of the task, but every change made to the task is only visible to themselves. They have to done their own work and change the _DONE_ status. The status update is visible to the task assigner.
   - Task can be read only or allowed to be modified by every worker and also the publisher. A read only task can only be modified by the publisher. Again, any operations of this should be recorded.
4. A calendar is available to every user to see the everyday task and the uncompleted tasks.
   - When clicking a specific date of tasks, a modal should be shown up to display tasks details.

**Dolphin** will **not** meet the following requirements:

1. Publish a group task to several subgroups of people.
2. Share the tasks to other people.
