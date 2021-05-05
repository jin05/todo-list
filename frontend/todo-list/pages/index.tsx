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

const getServerSideAuth = createGetServerSideAuth({
  pems: process.env.NODE_ENV === "production" ? pems : localPems,
});
const useAuth = createUseAuth({
  pems: process.env.NODE_ENV === "production" ? pems : localPems,
});

const Home = (props: { initialAuth: AuthTokens }) => {
  const auth = useAuth(props.initialAuth);
  const { login, logout } = useAuthFunctions();

  useEffect(() => {
    if (auth) {
      fetch(process.env.API_HOST + "/user", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: auth.idToken,
        },
      }).then((res) => {
        res.json().then((value) => {
          console.log("JSON: " + value);
        });
      });
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
  return { props: { initialAuth } };
};

export default Home;
