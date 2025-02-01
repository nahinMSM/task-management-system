import React from 'react';
import { useQuery } from '@tanstack/react-query';

interface Task {
  id: number;
  title: string;
  completed: boolean;
}

const getTasks = async (): Promise<Task[]> => {
  const response = await fetch('/api/tasks');
  const data = await response.json();
  return data;
};

const TaskList: React.FC = () => {
  const { data: tasks, isLoading, isError, error } = useQuery<Task[], Error>({
    queryKey: ['tasks'],
    queryFn: getTasks,
  });

  if (isLoading) {
    return <div>Carregando...</div>;
  }

  if (isError) {
    return <div>Erro: {error?.message}</div>;
  }

  return (
    <div>
      <h1>Lista de Tarefas</h1>
      <ul>
        {tasks?.map((task) => (
          <li key={task.id}>
            {task.title} {task.completed ? '(Concluída)' : '(Pendente)'}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TaskList;
