import cn from 'classnames'
import React, { type ButtonHTMLAttributes } from 'react'
import styles from './styles.module.scss'

type Props = ButtonHTMLAttributes<HTMLButtonElement> & {
  icon?: React.ReactNode;
  losePaddingLeft?: boolean;
  losePaddingRight?: boolean;
  variant?: 'primary' | 'secondary' | 'plain';
  size?: 'base' | 'lg' | 'md';
  iconOnly?: boolean;
}

const ButtonBase = ({ icon, losePaddingLeft = false, losePaddingRight = false, variant = 'secondary', size = 'base', iconOnly = false, className, ...props }: Props) => {
  return (
    <button className={cn(styles.button, styles[`button__${variant}`], styles[`button__${size}`], className, {
      [styles.button__losePaddingLeft]: losePaddingLeft,
      [styles.button__losePaddingRight]: losePaddingRight,
      [styles.button__iconOnly]: iconOnly
    })} {...props}>
      {icon && <span className={styles.button_icon}>{icon}</span>}

      <span className={styles.button_text}>
        {props.children}
      </span>
    </button>
  )
}

export default ButtonBase
