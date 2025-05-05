import { Flex } from '@chakra-ui/react'
import SVG from 'react-inlinesvg'
import ButtonBase from '../../components/button-base'

const NodeActions = () => {
  return (
    <Flex flexDirection="row" gap={'12px'}>
      <ButtonBase variant={'primary'} icon={<SVG src='/icons/ic_24_coins.svg' />}>
        Claim EAI
      </ButtonBase>

      <ButtonBase losePaddingLeft icon={<SVG src='/icons/ic_24_pause_circle.svg' />}>
        Stop service
      </ButtonBase>

      <ButtonBase icon={<SVG src='/icons/ic_24_clear.svg' />}>
        Reset setting
      </ButtonBase>
    </Flex>
  )
}

export default NodeActions
