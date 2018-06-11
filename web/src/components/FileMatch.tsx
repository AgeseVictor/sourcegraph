import React from 'react'
import { Link } from 'react-router-dom'
import * as GQL from '../backend/graphqlschema'
import { SymbolIcon } from '../symbols/SymbolIcon'
import { pluralize } from '../util/strings'
import { toPositionOrRangeHash } from '../util/url'
import { CodeExcerpt } from './CodeExcerpt'
import { CodeExcerpt2 } from './CodeExcerpt2'
import { RepoFileLink } from './RepoFileLink'
import { Props as ResultContainerProps, ResultContainer } from './ResultContainer'

export type IFileMatch = Partial<Pick<GQL.IFileMatch, 'symbols' | 'limitHit'>> & {
    file: Pick<GQL.IFile, 'path' | 'url'> & { commit: Pick<GQL.IGitCommit, 'oid'> }
    repository: Pick<GQL.IRepository, 'name' | 'url'>
    lineMatches: ILineMatch[]
}

export type ILineMatch = Pick<GQL.ILineMatch, 'preview' | 'lineNumber' | 'offsetAndLengths' | 'limitHit'>

interface IMatchItem {
    highlightRanges: {
        start: number
        highlightLength: number
    }[]
    preview: string
    line: number
    repoName: string
    repoURL: string
    filePath: string
    fileURL: string
    commitID: string
}

interface Props {
    /**
     * The file match search result.
     */
    result: IFileMatch

    /**
     * The icon to show left to the title.
     */
    icon: React.ComponentType<{ className?: string }>

    /**
     * Called when the file's search result is selected.
     */
    onSelect: () => void

    /**
     * Whether this file should be rendered as expanded.
     */
    expanded: boolean

    /**
     * Whether or not to show all matches for this file, or only a subset.
     */
    showAllMatches: boolean

    isLightTheme: boolean

    allExpanded?: boolean
}

const subsetMatches = 2

// Dev flag for disabling syntax highlighting on search results pages.
const NO_SEARCH_HIGHLIGHTING = localStorage.getItem('noSearchHighlighting') !== null

export class FileMatch extends React.PureComponent<Props> {
    public render(): React.ReactNode {
        const result = this.props.result
        const items: IMatchItem[] = this.props.result.lineMatches.map(m => ({
            highlightRanges: m.offsetAndLengths.map(offsetAndLength => ({
                start: offsetAndLength[0],
                highlightLength: offsetAndLength[1],
            })),
            preview: m.preview,
            line: m.lineNumber,
            repoName: result.repository.name,
            repoURL: result.repository.url,
            filePath: result.file.path,
            fileURL: result.file.url,
            commitID: result.file.commit.oid,
        }))

        const title = (
            <RepoFileLink
                repoPath={result.repository.name}
                repoURL={result.repository.url}
                filePath={result.file.path}
                fileURL={result.file.url}
            />
        )

        let containerProps: ResultContainerProps

        const expandedChildren = this.getChildren(items, result.file.url, true)
        if (this.props.showAllMatches) {
            containerProps = {
                collapsible: true,
                defaultExpanded: this.props.expanded,
                icon: this.props.icon,
                title,
                expandedChildren,
                allExpanded: this.props.allExpanded,
            }
        } else {
            const len = items.length - subsetMatches
            containerProps = {
                collapsible: items.length > subsetMatches,
                defaultExpanded: this.props.expanded,
                icon: this.props.icon,
                title,
                collapsedChildren: this.getChildren(items, result.file.url, false),
                expandedChildren,
                collapseLabel: `Hide ${len} matches`,
                expandLabel: `Show ${len} more ${pluralize('match', len, 'matches')}`,
                allExpanded: this.props.allExpanded,
            }
        }

        return <ResultContainer {...containerProps} />
    }

    // If this grows any larger, it needs to be factored out into it's own component
    private getChildren = (items: IMatchItem[], fileURL: string, allMatches: boolean) => {
        const showItems = items
            .sort((a, b) => {
                if (a.line < b.line) {
                    return -1
                }
                if (a.line === b.line) {
                    if (a.highlightRanges[0].start < b.highlightRanges[0].start) {
                        return -1
                    }
                    if (a.highlightRanges[0].start === b.highlightRanges[0].start) {
                        return 0
                    }
                    return 1
                }
                return 1
            })
            .filter((item, i) => allMatches || i < subsetMatches)

        if (NO_SEARCH_HIGHLIGHTING) {
            return <CodeExcerpt2 urlWithoutPosition={fileURL} items={showItems} onSelect={this.props.onSelect} />
        }

        return (
            <div className="file-match__list">
                {/* Symbols */}
                {(this.props.result.symbols || []).map(symbol => (
                    <Link
                        to={symbol.url}
                        className="file-match__item"
                        key={`symbol:${symbol.name}${symbol.containerName}${symbol.url}`}
                    >
                        <SymbolIcon kind={symbol.kind} className="icon-inline mr-1" />
                        <code>
                            {symbol.name}{' '}
                            {symbol.containerName && <span className="text-muted">{symbol.containerName}</span>}
                        </code>
                    </Link>
                ))}
                {showItems.map((item, i) => {
                    const position = { line: item.line + 1, character: item.highlightRanges[0].start + 1 }
                    return (
                        <Link
                            to={`${item.fileURL}${toPositionOrRangeHash({ position })}`}
                            key={`linematch:${item.fileURL}${position.line}:${position.character}`}
                            className="file-match__item file-match__item-clickable"
                            onClick={this.props.onSelect}
                        >
                            <CodeExcerpt
                                repoPath={item.repoName}
                                commitID={item.commitID}
                                filePath={item.filePath}
                                previewWindowExtraLines={1}
                                highlightRanges={item.highlightRanges}
                                line={item.line}
                                className="file-match__item-code-excerpt"
                                isLightTheme={this.props.isLightTheme}
                            />
                        </Link>
                    )
                })}
            </div>
        )
    }
}
