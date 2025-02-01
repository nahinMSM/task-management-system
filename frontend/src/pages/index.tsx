import { GetServerSideProps } from 'next';
import axios from 'axios';
import { useState } from 'react';

interface Task {
  id: number;
  title: string;
  description: string;
  completed: boolean;
  createdAt: string;
}

interface Props {
  tasks: Task[];
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const query = context.query.q || '';
  const response = await axios.get(`http://localhost:8080/tasks?q=${query}`);
  return {
    props: {
      tasks: response.data,
    },
  };
};

const Home = ({ tasks: initialTasks }: Props) => {
  const [tasks, setTasks] = useState(initialTasks);
  const [newTask, setNewTask] = useState({ title: '', description: '' });
  const [editingTask, setEditingTask] = useState<Task | null>(null);

  const createTask = async () => {
    const response = await axios.post('http://localhost:8080/tasks', newTask);
    setTasks([...tasks, { ...newTask, id: response.data.id, completed: false, createdAt: new Date().toISOString() }]);
    setNewTask({ title: '', description: '' });
  };

  const updateTask = async (task: Task) => {
    await axios.put(`http://localhost:8080/tasks?id=${task.id}`, task);
    setTasks(tasks.map((t) => (t.id === task.id ? task : t)));
    setEditingTask(null);
  };

  const deleteTask = async (id: number) => {
    await axios.delete(`http://localhost:8080/tasks?id=${id}`);
    setTasks(tasks.filter((t) => t.id !== id));
  };

  return (
    <div className="container mx-auto p-6 max-w-2xl">
      <h1 className="text-3xl font-bold mb-6 text-center">Gerenciador de Tarefas</h1>

      {/* Formulário para criar ou editar tarefa */}
      <div className="bg-white p-4 rounded-lg shadow-md mb-6">
        <input
          type="text"
          placeholder="Título"
          value={editingTask ? editingTask.title : newTask.title}
          onChange={(e) =>
            editingTask
              ? setEditingTask({ ...editingTask, title: e.target.value })
              : setNewTask({ ...newTask, title: e.target.value })
          }
          className="border p-2 w-full rounded mb-2"
        />
        <input
          type="text"
          placeholder="Descrição"
          value={editingTask ? editingTask.description : newTask.description}
          onChange={(e) =>
            editingTask
              ? setEditingTask({ ...editingTask, description: e.target.value })
              : setNewTask({ ...newTask, description: e.target.value })
          }
          className="border p-2 w-full rounded mb-4"
        />
        {editingTask ? (
          <button onClick={() => updateTask(editingTask)} className="bg-green-500 text-white px-4 py-2 w-full rounded">
            Atualizar
          </button>
        ) : (
          <button onClick={createTask} className="bg-blue-500 text-white px-4 py-2 w-full rounded">
            Criar
          </button>
        )}
      </div>

      {/* Lista de tarefas */}
      <ul className="space-y-4">
        {tasks.map((task) => (
          <li key={task.id} className={`p-4 border rounded-lg shadow-md flex justify-between items-center ${task.completed ? 'bg-gray-100' : 'bg-white'}`}>
            <div>
              <h2 className="text-xl font-semibold">{task.title}</h2>
              <p className="text-gray-600">{task.description}</p>
            </div>
            <div className="flex space-x-2">
              <button
                onClick={() => setEditingTask(task)}
                className="bg-yellow-500 text-white px-4 py-2 rounded"
              >
                Editar
              </button>
              <button
                onClick={() => deleteTask(task.id)}
                className="bg-red-500 text-white px-4 py-2 rounded"
              >
                Excluir
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Home;