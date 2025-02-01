import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import TaskList from "../components/TaskList";
import '@testing-library/jest-dom';

const queryClient = new QueryClient();

jest.mock("../components/TaskList", () => ({
  ...jest.requireActual("../components/TaskList"),
  getTasks: jest.fn(() =>
    Promise.resolve([
      { id: 1, title: "Tarefa 1", completed: false },
      { id: 2, title: "Tarefa 2", completed: true },
    ])
  ),
}));

describe("TaskList", () => {
  it("deve renderizar a lista de tarefas corretamente", async () => {
    render(
      <QueryClientProvider client={queryClient}>
        <TaskList />
      </QueryClientProvider>
    );

    expect(await screen.findByText(/tarefa 1/i)).toBeInTheDocument();
    expect(await screen.findByText(/tarefa 2/i)).toBeInTheDocument();
  });
});
