import React, { HTMLAttributes } from 'react'

import classNames from 'classnames'
import type { MdiReactIconProps } from 'mdi-react'

import styles from './RepoRevision.module.scss'
import { mdiChevronDown } from "@mdi/js";
import { Icon } from "@sourcegraph/wildcard";

type RepoRevisionProps = HTMLAttributes<HTMLDivElement>

export const RepoRevisionWrapper: React.FunctionComponent<React.PropsWithChildren<RepoRevisionProps>> = ({
    children,
    className,
    ...rest
}) => (
    <div className={classNames(styles.repoRevisionContainer, className)} {...rest}>
        {children}
    </div>
)

export const RepoRevisionChevronDownIcon: React.FunctionComponent<React.PropsWithChildren<MdiReactIconProps>> = ({
    className,
    ...rest
}) => <Icon className={classNames(styles.breadcrumbIcon, className)} {...rest} svgPath={mdiChevronDown} inline={false} aria-hidden={true} />
