import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
} from "@material-ui/core";
import { FormEvent } from "react";
import { Todo } from "../pages/list";
import styled from "styled-components";
import { CreateInput, todoApiRequest, UpdateInput } from "../utils/apiHelper";

type Props = {
  open: boolean;
  onChange: (todo: Todo) => void;
  onClose: () => void;
  todo: Todo | null;
};

export const TodoDialog = (props: Props) => {
  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const title = (event.currentTarget.elements.namedItem(
      "title"
    ) as HTMLInputElement).value;
    const content = (event.currentTarget.elements.namedItem(
      "content"
    ) as HTMLInputElement).value;

    let method;
    let input: CreateInput | UpdateInput;
    if (props.todo) {
      method = "PUT";
      input = {
        todoID: props.todo.TodoID,
        title: title,
        content: content,
        checked: props.todo.Checked,
      } as UpdateInput;
    } else {
      method = "POST";
      input = { title: title, content: content } as CreateInput;
    }

    todoApiRequest(method, input).then((json) => {
      if (!json) return;
      const todo = JSON.parse(json) as Todo;
      props.onChange(todo);
    });
  };

  return (
    <Dialog open={props.open} onClose={props.onClose}>
      <DialogTitle>Todo</DialogTitle>
      <form onSubmit={handleSubmit}>
        <DialogContent>
          <Text
            name="title"
            margin="dense"
            fullWidth
            label="Title"
            defaultValue={props.todo?.Title}
          />
          <Text
            name="content"
            margin="dense"
            fullWidth
            multiline
            label="Content"
            defaultValue={props.todo?.Content}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={props.onClose}>cancel</Button>
          <Button type="submit">ok</Button>
        </DialogActions>
      </form>
    </Dialog>
  );
};

const Text = styled(TextField)`
  margin: 10px;
`;
