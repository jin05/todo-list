import { GetServerSideProps } from "next";
import styled from "styled-components";
import MuiList from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Checkbox from "@material-ui/core/Checkbox";
import React, { useState } from "react";
import {
  Button,
  DialogActions,
  DialogContent,
  Fade,
  InputBase,
  ListItemIcon,
  ListItemSecondaryAction,
  MenuItem,
  Paper,
  Popper,
  Select,
} from "@material-ui/core";
import SearchIcon from "@material-ui/icons/Search";
import IconButton from "@material-ui/core/IconButton";
import AddIcon from "@material-ui/icons/Add";
import DeleteIcon from "@material-ui/icons/Delete";
import { TodoDialog } from "../../components/TodoDialog";
import { getIdToken } from "../../utils/cookieHelper";
import {
  DeleteInput,
  todoApiRequest,
  UpdateInput,
} from "../../utils/apiHelper";
import { useRouter } from "next/router";
import Typography from "@material-ui/core/Typography";

export type Todo = {
  TodoID: number;
  UserID: number;
  Title: string;
  Content: string;
  Checked: boolean;
};

const List = (props: { todoList: Todo[] }) => {
  const router = useRouter();

  const [dialogOpen, setDialogOpen] = useState(false);
  const [todoList, setTodoList] = useState<Todo[]>(props.todoList || []);
  const [target, setTarget] = useState<Todo | null>(null);
  const [deleteTarget, setDeleteTarget] = useState<Todo | null>(null);
  const [
    deleteButtonEl,
    setDeleteButtonEl,
  ] = React.useState<HTMLButtonElement | null>(null);

  const handleCloseDialog = () => {
    setTarget(null);
    setDialogOpen(false);
  };

  const handleChangeTodo = (todo: Todo) => {
    const index = todoList.findIndex((t) => {
      return t.TodoID === todo.TodoID;
    });

    if (index < 0) {
      todoList.push(todo);
    } else {
      todoList.splice(index, 1, todo);
    }

    setTodoList(todoList);
    handleCloseDialog();
  };

  const handleClickItem = (todo: Todo) => (
    event: React.MouseEvent<HTMLDivElement, MouseEvent>
  ) => {
    if ((event.target as HTMLElement).tagName === "INPUT") {
      return;
    }

    setTarget(todo);
    setDialogOpen(true);
  };

  const handleChangeCheck = (todo: Todo) => (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const input = {
      todoID: todo.TodoID,
      title: todo.Title,
      content: todo.Content,
      checked: event.target.checked,
    } as UpdateInput;
    todoApiRequest("PUT", input);
  };

  const handleClickDeleteIcon = (todo: Todo) => (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    setDeleteTarget(todo);
    setDeleteButtonEl(event.currentTarget);
  };

  const handleClickDelete = () => {
    setDeleteButtonEl(null);
    if (!deleteTarget) return;
    todoApiRequest("DELETE", {
      todoID: deleteTarget.TodoID,
    } as DeleteInput).then(() => {
      router.replace(router.asPath);
    });
  };

  const handleSubmitSearch = (event: React.FormEvent<HTMLDivElement>) => {
    event.preventDefault();

    const formEl = event.target as HTMLFormElement;
    const keyword = formEl.elements.namedItem("keyword") as HTMLInputElement;
    const searchTarget = formEl.elements.namedItem(
      "searchTarget"
    ) as HTMLSelectElement;

    router.push({
      pathname: router.pathname,
      query: { keyword: keyword.value, searchTarget: searchTarget.value },
    });
  };

  return (
    <Contents>
      <ListPaper>
        <Typography align="center" variant="h3">
          TODO List
        </Typography>
        <OperationArea>
          <IconButton color="primary" onClick={() => setDialogOpen(true)}>
            <AddIcon />
          </IconButton>
          <TodoDialog
            open={dialogOpen}
            onChange={handleChangeTodo}
            onClose={handleCloseDialog}
            todo={target}
          />
          <SearchForm
            variant="outlined"
            component="form"
            onSubmit={handleSubmitSearch}
          >
            <IconButton type="submit">
              <SearchIcon />
            </IconButton>
            <SearchInput name="keyword" placeholder="Search" />
            <Select name="searchTarget" defaultValue="title">
              <MenuItem value="title">Title</MenuItem>
              <MenuItem value="content">Content</MenuItem>
            </Select>
          </SearchForm>
        </OperationArea>
        <MuiList>
          {todoList.map((todo) => {
            return (
              <ListItem
                key={todo.TodoID}
                button
                onClick={handleClickItem(todo)}
              >
                <ListItemIcon>
                  <Checkbox
                    defaultChecked={todo.Checked}
                    onChange={handleChangeCheck(todo)}
                  />
                </ListItemIcon>
                <ListItemText primary={todo.Title} />
                <ListItemSecondaryAction>
                  <IconButton onClick={handleClickDeleteIcon(todo)}>
                    <DeleteIcon />
                  </IconButton>
                  <Popper
                    open={deleteButtonEl !== null}
                    anchorEl={deleteButtonEl}
                    placement="top"
                    transition
                  >
                    {({ TransitionProps }) => (
                      <Fade {...TransitionProps} timeout={350}>
                        <Paper>
                          <DialogContent>削除しますか？</DialogContent>
                          <DialogActions>
                            <Button
                              color="primary"
                              onClick={() => setDeleteButtonEl(null)}
                            >
                              cancel
                            </Button>
                            <Button color="primary" onClick={handleClickDelete}>
                              ok
                            </Button>
                          </DialogActions>
                        </Paper>
                      </Fade>
                    )}
                  </Popper>
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
  const keyword = ctx.query.keyword || "";
  const searchTarget = ctx.query.searchTarget || "";

  const query_params = new URLSearchParams({
    keyword: keyword.toString(),
    searchTarget: searchTarget.toString(),
  });
  const res = await fetch(`${process.env.API_HOST}/todo/list?${query_params}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: idToken,
    },
    mode: "cors",
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

const SearchInput = styled(InputBase)``;
