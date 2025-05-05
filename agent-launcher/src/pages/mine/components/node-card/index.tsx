import type { Node } from '../../../../types/data';
import { formatNumber } from '../../../../utils/data';
import { useNodes } from '../../stores/useNodes';
import CardBase from '../card-base';
import styles from './styles.module.scss';

type Props = Node

const UNSELECTED_STYLE = {
  borderColor: '#EFEFEF',
  backgroundColor: '#FAFAFA',
}

const SELECTED_STYLE = {
  borderColor: 'linear-gradient(to right, #00F5A0 0%, #00D9F5 100%)',
  backgroundColor: '#fff',
  borderWidth: '2px',
}

const NodeCard = (props: Props) => {
  const { selectedItem: selectedNode, setSelectedItem: setSelectedNode } = useNodes();

  return (
    <CardBase
      className={styles.card}
      onClick={() => setSelectedNode(props)}
      {...(selectedNode?.id === props.id ? SELECTED_STYLE : UNSELECTED_STYLE)}
    >
      <h5 className={styles.card_name}>{props.name}</h5>

      <div className={styles.card_imageWrapper} style={{ background: `url(${props.image}) 50% / cover no-repeat` }}>
      </div>

      <h4 className={styles.card_earnings}>{formatNumber(props.earnings)} EAI</h4>

      <p className={styles.card_status}>{props.status}</p>
    </CardBase>
  )
}

export default NodeCard
