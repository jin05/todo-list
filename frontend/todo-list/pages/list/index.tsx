import { GetServerSideProps } from "next";
import styled from "styled-components";
import MuiList from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Checkbox from "@material-ui/core/Checkbox";
import React, { useState } from "react";
import {
  InputBase,
  ListItemIcon,
  ListItemSecondaryAction,
  Paper,
} from "@material-ui/core";
import SearchIcon from "@material-ui/icons/Search";
import IconButton from "@material-ui/core/IconButton";
import AddIcon from "@material-ui/icons/Add";
import DeleteIcon from "@material-ui/icons/Delete";
import { TodoDialog } from "../../components/TodoDialog";
import { getIdToken } from "../../utils/cookieHelper";

export type Todo = {
  TodoID: number;
  UserID: number;
  Title: string;
  Content: string;
  Checked: boolean;
};

const List = (props: { todoList: Todo[] }) => {
  const [dialogOpen, setDialogOpen] = useState(false);
  const [todoList, setTodoList] = useState(props.todoList);
  const [target, setTarget] = useState<Todo | null>(null);

  const handleDialogClose = () => {
    setTarget(null);
    setDialogOpen(false);
  };

  const handleTodoChange = (todo: Todo) => {
    const index = todoList.findIndex((t) => {
      return t.TodoID === todo.TodoID;
    });

    if (index < 0) {
      todoList.push(todo);
    } else {
      todoList.splice(index, 1, todo);
    }

    setTodoList(todoList);
    handleDialogClose();
  };

  const handleTodoItemClick = (todo: Todo) => {
    setTarget(todo);
    setDialogOpen(true);
  };

  return (
    <Contents>
      <ListPaper>
        <OperationArea>
          <IconButton color="primary" onClick={() => setDialogOpen(true)}>
            <AddIcon />
          </IconButton>
          <TodoDialog
            open={dialogOpen}
            onChange={handleTodoChange}
            onClose={handleDialogClose}
            todo={target}
          />
          <SearchForm variant="outlined" component="form">
            <IconButton type="submit">
              <SearchIcon />
            </IconButton>
            <SearchInput placeholder="Search" />
          </SearchForm>
        </OperationArea>
        <MuiList>
          {props.todoList.map((todo) => {
            return (
              <ListItem
                key={todo.TodoID}
                button
                onClick={() => {
                  handleTodoItemClick(todo);
                }}
              >
                <ListItemIcon>
                  <Checkbox checked={todo.Checked} />
                </ListItemIcon>
                <ListItemText primary={todo.Title} />
                <ListItemSecondaryAction>
                  <IconButton>
                    <DeleteIcon />
                  </IconButton>
                </ListItemSecondaryAction>
              </ListItem>
            );
          })}
        </MuiList>
      </ListPaper>
    </Contents>
  );
};

export const getServerSideProps: GetServerSideProps<{
  todoList: Todo[];
}> = async (ctx) => {
  const idToken = getIdToken(ctx);

  const res = await fetch(process.env.API_HOST + "/todo/list", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: idToken,
    },
  });

  let todoList: Todo[] = [];
  if (res.ok) {
    todoList = (await res.json()) as Todo[];
  } else {
    console.error(`${res.status}:${res.statusText}`);
    ctx.res.writeHead(302, { Location: "/" });
    ctx.res.end();
  }

  return { props: { todoList: todoList } };
};

export default List;

const Contents = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 50px;
`;

const ListPaper = styled(Paper)`
  width: 600px;
`;

const OperationArea = styled.div`
  display: flex;
  margin: 10px;
`;

const SearchForm = styled(Paper)`
  margin-left: auto;
  width: 300px;
`;

const SearchInput = styled(InputBase)`
  width: 200px;
`;
