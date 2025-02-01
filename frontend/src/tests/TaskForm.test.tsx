import { render, screen, fireEvent } from "@testing-library/react";
import TaskForm from "../components/TaskForm";

describe("TaskForm", () => {
  it("deve chamar a função de callback após salvar a tarefa", async () => {
    const mockOnTaskSaved = jest.fn();
    render(<TaskForm onTaskSaved={mockOnTaskSaved} />);

    fireEvent.change(screen.getByPlaceholderText(/título/i), {
      target: { value: "Nova Tarefa" },
    });

    fireEvent.change(screen.getByPlaceholderText(/descrição/i), {
      target: { value: "Descrição da nova tarefa" },
    });

    fireEvent.click(screen.getByText(/criar/i));

    await new Promise((resolve) => setTimeout(resolve, 0));

    expect(mockOnTaskSaved).toHaveBeenCalled();
  });

  it("não deve chamar onTaskSaved se os campos estiverem vazios", async () => {
    const mockOnTaskSaved = jest.fn();
    render(<TaskForm onTaskSaved={mockOnTaskSaved} />);

    fireEvent.click(screen.getByText(/criar/i));

    await new Promise((resolve) => setTimeout(resolve, 0));

    expect(mockOnTaskSaved).not.toHaveBeenCalled();
  });
});