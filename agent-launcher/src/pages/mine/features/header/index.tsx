import { Box, Flex } from '@chakra-ui/react'
import SVG from 'react-inlinesvg'
import ButtonBase from '../../components/button-base'
import styles from './styles.module.scss'
import { copyToClipboard } from '../../../../utils/extension'

const MineHeader = () => {
  return (
    <Flex width={'100%'} justifyContent={'space-between'} className={styles.header}>
      <Box className={styles.header_left}>
        <Flex alignItems={'center'} gap={'8px'} marginBottom={'4px'} className={styles.header_left_address}>
          <p>0x223...2435</p>
          <SVG src='/icons/ic_24_copy.svg' width={24} height={24} onClick={() => copyToClipboard('0x223...2435')} className={styles.header_left_address_icon} />
        </Flex>

        <Flex alignItems={'center'} gap={'8px'} className={styles.header_left_balance}>
          <h4>13,343.34</h4>
          <span className={styles.header_left_balance_unit}>EAI</span>
        </Flex>
      </Box>

      <Flex gap={'10px'}>
        <ButtonBase variant={'primary'} size={'md'}>
          Deposit
        </ButtonBase>
        <ButtonBase variant='plain' iconOnly size={'md'}>
          <SVG src='/icons/ic_44_dots-more_circle.svg' />
        </ButtonBase>
      </Flex>
    </Flex>
  )
}

export default MineHeader
