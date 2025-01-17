import InformationOutlineIcon from 'mdi-react/InformationOutlineIcon'

import { formatSearchParameters } from '@sourcegraph/common'
import { SyntaxHighlightedSearchQuery } from '@sourcegraph/search-ui'
import { AggregateStreamingSearchResults } from '@sourcegraph/shared/src/search/stream'
import { Link, H3, createLinkUrl, Tooltip, Icon } from '@sourcegraph/wildcard'

import styles from './QuerySuggestion.module.scss'

interface LuckySearchProps {
    alert: Required<AggregateStreamingSearchResults>['alert'] | undefined
}

export const LuckySearch: React.FunctionComponent<React.PropsWithChildren<LuckySearchProps>> = ({ alert }) =>
    alert?.kind && alert.kind !== 'lucky-search-queries' ? null : (
        <div className={styles.root}>
            <H3>
                Also showing results for:
                <Tooltip content="We returned all the results for your query. We also added results you might be interested in for similar queries. Below are similar queries we ran.">
                    <Icon
                        size="sm"
                        className="ml-1"
                        as={InformationOutlineIcon}
                        tabIndex={0}
                        aria-label="More information"
                    />
                </Tooltip>
            </H3>
            <ul className={styles.container}>
                {alert?.proposedQueries?.map(entry => (
                    <li className="mt-2" key={entry.query}>
                        <Link
                            to={createLinkUrl({
                                pathname: '/search',
                                search: formatSearchParameters(new URLSearchParams({ q: entry.query })),
                            })}
                        >
                            <span className={styles.suggestion}>
                                <SyntaxHighlightedSearchQuery query={entry.query} />
                            </span>
                            <i>{`— ${entry.description}`}</i>
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    )
