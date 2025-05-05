import { Box, Flex } from "@chakra-ui/react";
import { PropsWithChildren } from "react";
import Header from "../header";
import s from "./styles.module.scss";

type Props = {};

const MainLayout = (props: Props & PropsWithChildren) => {
  const { children } = props;
  return (
    <Flex className={s.container}>
      <Header />
      <Box className={s.content}>{children}</Box>
    </Flex>
  );
};

export default MainLayout;
