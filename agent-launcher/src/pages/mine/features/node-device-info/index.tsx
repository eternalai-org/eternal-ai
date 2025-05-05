import { Flex } from '@chakra-ui/react';
import { useMemo } from 'react';
import ConfigInfoCard from '../../components/config-info-card';
import { useNodes } from '../../stores/useNodes';
import styles from './styles.module.scss';

const NodeDeviceInfoItem = ({ label, value }: { label: string, value: string | number }) => {
  return (
    <Flex flexDirection="column" alignItems="center" flex={1}>
      <h4 className={styles.value}>{value}</h4>
      <p className={styles.label}>{label}</p>
    </Flex>
  )
}

const NodeDeviceInfo = () => {
  const selectedNode = useNodes(state => state.selectedItem);

  const values = useMemo(() => {
    if (!selectedNode) return [];

    return [
      { label: "DEVICE", value: selectedNode.device.name },
      { label: "OS", value: selectedNode.device.os },
      { label: "PROCESSOR", value: selectedNode.device.processor },
      { label: "RAM", value: selectedNode.device.ram + " GB" },
      { label: "GPU", value: selectedNode.device.gpu },
      { label: "GPU CORE", value: selectedNode.device.gpu_cores + "-Core" },
    ]
  }, [selectedNode])

  return (
    <ConfigInfoCard configName="Device information">
      <Flex flexDirection="row" width={'100%'} justifyContent={'space-between'}>
        {values.map((item) => (
          <NodeDeviceInfoItem key={item.label} {...item} />
        ))}
      </Flex>
    </ConfigInfoCard>
  )
}

export default NodeDeviceInfo
