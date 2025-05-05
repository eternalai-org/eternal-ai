import type { PropsWithChildren } from 'react';
import CardBase from '../card-base';
import styles from './styles.module.scss';
type Props = PropsWithChildren & {
  configName: string;
}

const ConfigInfoCard = ({ configName, children }: Props) => {
  return (
    <CardBase className={styles.card} padding={'0'}>
      <div className={styles.card_header}>
        <h5 className={styles.card_header_name}>{configName}</h5>
      </div>

      <div className={styles.card_content}>
        {children}
      </div>
    </CardBase>
  )
}

export default ConfigInfoCard
