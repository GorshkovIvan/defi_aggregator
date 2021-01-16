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
                <Table.Cell>{ result.pair }</Table.Cell>
                <Table.Cell>{ result.amount }</Table.Cell>
                <Table.Cell>{ result.pool_sz }</Table.Cell>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment>
                <Header>Optimal Portfolio</Header>
                <Table striped>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Ranking</Table.HeaderCell>
                            <Table.HeaderCell>Pair</Table.HeaderCell>
                            <Table.HeaderCell>Amount</Table.HeaderCell>
                            <Table.HeaderCell>Pool Size</Table.HeaderCell>
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