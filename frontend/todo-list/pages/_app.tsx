import React, { useEffect } from "react";
import { AppProps } from "next/app";
import Layout from "../components/Layout";
import { Amplify } from "@aws-amplify/core";
import { Auth } from "@aws-amplify/auth";
import Head from "next/head";

Amplify.configure({
  Auth: {
    region: "ap-northeast-1", //! Konfiguration
    userPoolId: process.env.USER_POOL_ID,
    userPoolWebClientId: process.env.USER_POOL_CLIENT_ID,

    // OPTIONAL - Configuration for cookie storage
    // Note: if the secure flag is set to true, then the cookie transmission requires a secure protocol
    // example taken from https://aws-amplify.github.io/docs/js/authentication
    cookieStorage: {
      // REQUIRED - Cookie domain (only required if cookieStorage is provided)
      // This should be the subdomain in production as the cookie should only
      // be present for the current site
      domain: process.env.AUTH_COOKIE_DOMAIN,
      // OPTIONAL - Cookie path
      path: "/",
      // OPTIONAL - Cookie expiration in days
      expires: 7,
      // OPTIONAL - Cookie secure flag
      // Either true or false, indicating if the cookie transmission requires a secure protocol (https).
      // The cookie can be secure in production
      secure: false,
    },
  },
});

Auth.configure({
  oauth: {
    domain: process.env.IDP_DOMAIN,
    scope: ["email", "openid"],
    // we need the /autologin step in between to set the cookies properly,
    // we don't need that when signing out though
    redirectSignIn: process.env.REDIRECT_SIGN_IN,
    redirectSignOut: process.env.REDIRECT_SIGN_OUT,
    responseType: "token",
  },
});

const MyApp = ({ Component, pageProps }: AppProps): JSX.Element => {
  useEffect(() => {
    const jssStyles = document.querySelector("#jss-server-side");
    if (jssStyles?.parentElement) {
      jssStyles.parentElement.removeChild(jssStyles);
    }
  }, []);

  return (
    <Layout>
      <Head>
        <title>{"My Todo List"}</title>
        <meta charSet="utf-8" />
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      <Component {...pageProps} />
    </Layout>
  );
};

export default MyApp;
