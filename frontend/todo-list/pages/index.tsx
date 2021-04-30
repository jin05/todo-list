import React, { useEffect } from "react";
import { GetServerSideProps } from "next";
import {
  AuthTokens,
  createGetServerSideAuth,
  createUseAuth,
  useAuthFunctions,
} from "aws-cognito-next";

import pems from "../local_pems.json";

const getServerSideAuth = createGetServerSideAuth({ pems });
const useAuth = createUseAuth({ pems });

const Home = (props: { initialAuth: AuthTokens }) => {
  const auth = useAuth(props.initialAuth);
  const { login, logout } = useAuthFunctions();

  useEffect(() => {
    if (auth) {
      window.localStorage.setItem("auth", JSON.stringify(auth));
      return;
    }
  }, [auth]);

  return (
    <React.Fragment>
      {auth ? (
        <button type="button" onClick={() => logout()}>
          sign out
        </button>
      ) : (
        <React.Fragment>
          <button type="button" onClick={() => login()}>
            sign in
          </button>
        </React.Fragment>
      )}
    </React.Fragment>
  );
};

export const getServerSideProps: GetServerSideProps<{
  initialAuth: AuthTokens;
}> = async (context) => {
  const initialAuth = getServerSideAuth(context.req);
  console.log(initialAuth);
  if (initialAuth) {
    const res = await fetch("http://localhost:8080/user", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ token: initialAuth.idToken }),
    });
    console.log(await res.json());
  }
  return { props: { initialAuth } };
};

export default Home;
