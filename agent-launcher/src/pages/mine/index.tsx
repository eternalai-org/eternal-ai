import { useEffect } from 'react';
import MainLayout from "../../components/layout";
import { MOCK_NODES } from './__mock__/nodes';
import MineHeader from "./features/header";
import NodeActions from './features/node-actions';
import NodeConfigInfo from './features/node-config-info';
import NodeDeviceInfo from './features/node-device-info';
import NodeList from './features/node-list';
import { useNodes } from "./stores/useNodes";
import { Flex } from '@chakra-ui/react';

const Mine = () => {
  const setNodes = useNodes(state => state.setItems);

  useEffect(() => {
    setNodes(MOCK_NODES);
  }, []);

  return (
    <MainLayout>
      <Flex flexDirection="column" gap={'24px'} width={'100%'}>
        <MineHeader />
        <NodeList />
        <NodeConfigInfo />
        <NodeDeviceInfo />
        <NodeActions />
      </Flex>
    </MainLayout>
  );
};

export default Mine;
