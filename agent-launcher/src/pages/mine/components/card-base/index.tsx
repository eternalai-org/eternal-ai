import cn from 'classnames'
import { type HTMLAttributes } from 'react'
import styles from './styles.module.scss'

type Props = HTMLAttributes<HTMLDivElement> & {
  padding?: string;
  borderRadius?: string;
  borderWidth?: string;
  borderColor?: string;
  backgroundColor?: string;
}

const CardBase = ({ children, className, padding = '24px', borderRadius = '12px', borderWidth = '1px', borderColor = '#EFEFEF', backgroundColor = '#fff', ...props }: Props) => {
  return (
    <div className={styles.card} {...props} style={{
      '--real-padding': padding,
      '--border-radius': borderRadius,
      '--border-width': borderWidth,
      '--border-color': borderColor,
      '--background-color': backgroundColor,
    } as React.CSSProperties}>
      <div className={styles.card_inner}>
        <div className={cn(styles.card_inner_content, className)}>
          {children}
        </div>
      </div>
    </div>
  )
}

export default CardBase
