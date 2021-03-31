import React from 'react';
import { Table, Header, Segment, Label } from 'semantic-ui-react'
export default function ResultsTable({results}) {
    const rows = results.map(((result, index) => {
        let color;
        if (index === 0) {
            color='yellow';
        } else if (index === 1) {
            color='grey';
        } else if (index === 2) {
            color='orange';
        }
        return (
            <Table.Row key={ index }>
                <Table.Cell>
                    <Label ribbon color={color}>{ index + 1 }</Label>
                </Table.Cell>
                <Table.Cell>{ result.tokenorpair }</Table.Cell>
                <Table.Cell>{ result.pool}</Table.Cell>
                <Table.Cell>{ result.amount }</Table.Cell>
                <Table.Cell>{ result.percentageofportfolio }</Table.Cell>
                <Table.Cell>{ result.roi_estimate }</Table.Cell>
                <Table.Cell>{ result.risk_setting }</Table.Cell>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment>

                <Table striped>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Ranking</Table.HeaderCell>
                            <Table.HeaderCell>Token/Pair</Table.HeaderCell>
                            <Table.HeaderCell>Pool</Table.HeaderCell>
                            <Table.HeaderCell>Amount</Table.HeaderCell>
                            <Table.HeaderCell>% Portfolio</Table.HeaderCell>
                            <Table.HeaderCell>ROI Estimate</Table.HeaderCell>
                            <Table.HeaderCell>Risk setting</Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { rows }
                    </Table.Body>
                </Table>
            </Segment>
        </div>
    );
}

//                 <Header>Optimal Portfolio</Header>