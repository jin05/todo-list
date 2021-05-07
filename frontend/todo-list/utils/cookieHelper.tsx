import nookies, { parseCookies } from "nookies";
import { GetServerSidePropsContext } from "next";
import { ParsedUrlQuery } from "querystring";

export const getIdToken = (
  context?: GetServerSidePropsContext<ParsedUrlQuery>
): string => {
  let cookies;
  if (context) {
    cookies = nookies.get(context);
  } else {
    cookies = parseCookies();
  }

  const userID =
    cookies[
      `CognitoIdentityServiceProvider.${process.env.USER_POOL_CLIENT_ID}.LastAuthUser`
    ];
  return cookies[
    `CognitoIdentityServiceProvider.${process.env.USER_POOL_CLIENT_ID}.${userID}.idToken`
  ];
};
