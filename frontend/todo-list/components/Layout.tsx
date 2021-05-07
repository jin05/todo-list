import React, { ReactNode, useEffect, useState } from "react";
import styled, {
  ThemeProvider as StyledComponentsThemeProvider,
} from "styled-components";
import Theme from "./Theme";
import { ThemeProvider as MaterialUIThemeProvider } from "@material-ui/core/styles";

type Props = {
  children?: ReactNode;
};

const Layout = ({ children }: Props) => {
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) return <></>;

  return (
    <MaterialUIThemeProvider theme={Theme}>
      <StyledComponentsThemeProvider theme={Theme}>
        <Contents>{children}</Contents>
      </StyledComponentsThemeProvider>
    </MaterialUIThemeProvider>
  );
};

export default Layout;

const Contents = styled.div`
  width: 100%;
`;
