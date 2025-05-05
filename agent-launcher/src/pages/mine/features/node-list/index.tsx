import { Flex } from '@chakra-ui/react';
import SVG from 'react-inlinesvg';
import { useNodes } from '../../stores/useNodes';
import CardBase from '../../components/card-base';
import NodeCard from '../../components/node-card';
import styles from './styles.module.scss';

const NodeList = () => {
  const nodes = useNodes(state => state.items);

  return (
    <Flex flexDirection="row" gap={'24px'}>
      {nodes.map((node) => (
        <NodeCard key={node.id} {...node} />
      ))}

      <CardBase className={styles.newCard} borderColor='#000000'>
        <Flex width={'100%'} height={'100%'} alignItems={'center'} justifyContent={'center'} flexDirection={'column'}>
          <SVG src='/icons/ic_48_laptop-code.svg' width={48} height={48} className={styles.newCard_icon} />

          <p className={styles.newCard_text}>run a new node</p>
        </Flex>
      </CardBase>
    </Flex>
  )
}

export default NodeList
