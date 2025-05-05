import { Box, Flex } from '@chakra-ui/react';
import SVG from 'react-inlinesvg';
import ConfigInfoCard from '../../components/config-info-card';
import { useNodes } from '../../stores/useNodes';
import styles from './styles.module.scss';
import { formatNumber, truncateAddress } from '../../../../utils/data';
const NodeConfigInfo = () => {
  const selectedNode = useNodes(state => state.selectedItem);

  return (
    <Flex flexDirection="row" gap={'24px'} width={'100%'} className={styles.container}>
      <Box flex={1} height={'100%'} alignSelf={'stretch'}>
        <ConfigInfoCard configName="Network">
          <Flex alignItems={'center'} gap={'24px'} padding={'12px 24px 0 24px'} height={'128px'}>
            <Flex alignItems={'center'} justifyContent={'center'} width={'80px'} height={'80px'} style={{background: `url(${selectedNode?.network.image}) 50% / cover no-repeat`}}>
            </Flex>

            <Box>
              <h4>{selectedNode?.network.name}</h4>

              <Flex gap={'4px'}>
                <p className={`${styles.health_status} ${styles[`health_status__${selectedNode?.network.health_status.toLowerCase()}`]}`}>{selectedNode?.network.health_status}</p>
                <p className={styles.dot}>•</p>
                <p className={styles.progress_status}>{selectedNode?.network.progress_status}</p>
              </Flex>
            </Box>
          </Flex>
        </ConfigInfoCard>
      </Box>

      <Box flex={1} height={'100%'} alignSelf={'stretch'}>
        <ConfigInfoCard configName="Model information">
          <Flex alignItems={'center'} gap={'24px'} padding={'0 24px'} height={'128px'}>
            <Flex alignItems={'center'} justifyContent={'center'} width={'120px'} height={'120px'} borderRadius={'12px'} overflow={'cover'} style={{background: `url(${selectedNode?.model.image}) 50% / cover no-repeat`}}>
            </Flex>

            <Box>
              <h4>{selectedNode?.model.name}</h4>
              <p className={styles.memory}>{selectedNode?.model.memory} GB</p>
            </Box>
          </Flex>
        </ConfigInfoCard>
      </Box>

      <Box flex={1} height={'100%'} alignSelf={'stretch'} className={styles.node_onchain_data}>
        <ConfigInfoCard configName="Node onchain data">
          <Flex flexDirection={'column'} gap={'4px'} height={'128px'}>
            <Flex width={"100%"} padding={'8.5px 0'} justifyContent={'space-between'} alignItems={'center'}>
              <p>Address</p>

              <a className={styles.right} href={`https://etherscan.io/address/${selectedNode?.onchain_data.address}`} target="_blank" rel="noreferrer">
                <span className={styles.right_text}>{truncateAddress(selectedNode?.onchain_data.address || '')}</span>
                <span className={styles.right_icon}><SVG src='/icons/ic_20_arrow-top-right.svg' /></span>
              </a>
            </Flex>

            <Flex width={"100%"} padding={'8.5px 0'} justifyContent={'space-between'} alignItems={'center'}>
              <p>ID</p>
              <p className={styles.right}>{selectedNode?.onchain_data.id}</p>
            </Flex>

            <Flex width={"100%"} padding={'8.5px 0'} justifyContent={'space-between'} alignItems={'center'}>
              <p>Processing tasks</p>
              <a className={styles.right} href={`https://etherscan.io/address/${selectedNode?.onchain_data.address}`} target="_blank" rel="noreferrer">
                <span className={styles.right_text}>{formatNumber(selectedNode?.onchain_data.processing_tasks || 0)}</span>
                <span className={styles.right_icon}><SVG src='/icons/ic_20_arrow-top-right.svg' /></span>
              </a>
            </Flex>
          </Flex>
        </ConfigInfoCard>
      </Box>
    </Flex>
  )
}

export default NodeConfigInfo
