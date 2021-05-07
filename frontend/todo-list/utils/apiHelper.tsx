import { getIdToken } from "./cookieHelper";

export type CreateInput = {
  title: string;
  content: string;
};

export type UpdateInput = {
  todoID: number;
  title: string;
  content: string;
  checked: boolean;
};

export type DeleteInput = {
  todoID: number;
};

export const todoApiRequest = async (
  method: string,
  reqStruct: CreateInput | UpdateInput | DeleteInput | undefined
): Promise<string | null> => {
  const idToken = getIdToken();

  const params: RequestInit = {
    method: method,
    headers: {
      "Content-Type": "application/json",
      Authorization: idToken,
    },
    mode: "cors",
  };

  if (reqStruct) {
    params.body = JSON.stringify(reqStruct);
  }

  const res = await fetch(process.env.API_HOST + "/todo", params);
  if (res.ok) {
    return JSON.stringify(await res.json());
  } else {
    if (res.status === 403) {
      if (typeof window !== "undefined") {
        window.location.href = "/";
      }
    }
  }

  return null;
};
