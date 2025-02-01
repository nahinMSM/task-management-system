import axios from "../utils/api";

export const getTasks = async (query: string = "") => {
  try {
    const response = await axios.get(`/tasks?q=${query}`);
    return response.data;
  } catch (error) {
    console.error("Erro ao buscar tarefas:", error);
    throw new Error("Não foi possível buscar as tarefas.");
  }
};

export const createTask = async (task: { title: string; description: string }) => {
  try {
    const response = await axios.post("/tasks", task);
    return response.data;
  } catch (error) {
    console.error("Erro ao criar tarefa:", error);
    throw new Error("Não foi possível criar a tarefa.");
  }
};

export const updateTask = async (id: number, task: { title: string; description: string; completed: boolean }) => {
  try {
    const response = await axios.put(`/tasks/${id}`, task);
    return response.data;
  } catch (error) {
    console.error("Erro ao atualizar tarefa:", error);
    throw new Error("Não foi possível atualizar a tarefa.");
  }
};

export const deleteTask = async (id: number) => {
  try {
    const response = await axios.delete(`/tasks/${id}`);
    return response.data;
  } catch (error) {
    console.error("Erro ao excluir tarefa:", error);
    throw new Error("Não foi possível excluir a tarefa.");
  }
};
