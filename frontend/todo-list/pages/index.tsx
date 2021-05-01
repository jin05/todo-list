import React, { useEffect } from "react";
import { GetServerSideProps } from "next";
import {
  AuthTokens,
  createGetServerSideAuth,
  createUseAuth,
  useAuthFunctions,
} from "aws-cognito-next";

import pems from "../local_pems.json";
import localPems from "../local_pems.json";

const getServerSideAuth = createGetServerSideAuth({
  pems: process.env.NODE_ENV === "production" ? pems : localPems,
});
const useAuth = createUseAuth({
  pems: process.env.NODE_ENV === "production" ? pems : localPems,
});

const Home = (props: {
  initialAuth: AuthTokens;
  userJson: string | undefined;
}) => {
  const auth = useAuth(props.initialAuth);
  const { login, logout } = useAuthFunctions();

  useEffect(() => {
    console.log(props.userJson);
    if (props.userJson) {
      window.localStorage.setItem("user", props.userJson);
    }
  }, [props.userJson]);

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
  userJson: string | undefined;
}> = async (context) => {
  const initialAuth = getServerSideAuth(context.req);
  let userJson: string | undefined;
  if (initialAuth) {
    const res = await fetch(
      process.env.API_HOST || "http://localhost:8080/user",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: initialAuth.idToken,
        },
      }
    );
    userJson = await res.json();
  }
  return { props: { initialAuth, userJson } };
};

export default Home;
