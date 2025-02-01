import { useState } from "react";
import { createTask, updateTask } from "../services/taskService";

interface TaskFormProps {
  initialTask?: { id: number; title: string; description: string };
  onTaskSaved: () => void;
}

const TaskForm = ({ initialTask, onTaskSaved }: TaskFormProps) => {
  const [task, setTask] = useState(
    initialTask || { title: "", description: "" }
  );

  const handleSubmit = async () => {
    if (initialTask) {
      await updateTask(initialTask.id, { ...task, completed: false });
    } else {
      await createTask(task);
    }
    onTaskSaved();
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Título"
        value={task.title}
        onChange={(e) => setTask({ ...task, title: e.target.value })}
      />
      <input
        type="text"
        placeholder="Descrição"
        value={task.description}
        onChange={(e) => setTask({ ...task, description: e.target.value })}
      />
      <button onClick={handleSubmit}>{initialTask ? "Atualizar" : "Criar"}</button>
    </div>
  );
};

export default TaskForm;