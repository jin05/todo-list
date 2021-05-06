import React, { useEffect } from "react";
import { GetServerSideProps } from "next";
import {
  AuthTokens,
  createGetServerSideAuth,
  createUseAuth,
  useAuthFunctions,
} from "aws-cognito-next";

import pems from "../pems/pems.json";
import localPems from "../pems/local_pems.json";
import { useRouter } from "next/router";
import { Button } from "@material-ui/core";

const getServerSideAuth = createGetServerSideAuth({
  pems: process.env.NODE_ENV === "production" ? pems : localPems,
});
const useAuth = createUseAuth({
  pems: process.env.NODE_ENV === "production" ? pems : localPems,
});

const Home = (props: { initialAuth: AuthTokens; userJson: string | null }) => {
  const router = useRouter();
  const auth = useAuth(props.initialAuth);
  const { login, logout } = useAuthFunctions();

  useEffect(() => {
    if (!auth) return;
    if (props.userJson) {
      localStorage.setItem("user", props.userJson);
      router.push("/list").catch((reason) => console.warn(reason));
    }
  }, [auth]);

  return (
    <React.Fragment>
      {auth ? (
        <Button size="large" variant="outlined" onClick={logout}>
          sign out
        </Button>
      ) : (
        <React.Fragment>
          <Button size="large" variant="outlined" onClick={login}>
            sign in
          </Button>
        </React.Fragment>
      )}
    </React.Fragment>
  );
};

export const getServerSideProps: GetServerSideProps<{
  initialAuth: AuthTokens;
  userJson: string | null;
}> = async (context) => {
  const initialAuth = getServerSideAuth(context.req);
  let userJson = null;
  if (initialAuth) {
    const res = await fetch(process.env.API_HOST + "/user", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: initialAuth.idToken,
      },
    });
    if (res.ok) {
      userJson = JSON.stringify(await res.json());
    } else {
      console.log(res.statusText);
    }
  }
  return { props: { initialAuth, userJson } };
};

export default Home;
