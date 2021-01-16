import React from 'react';
import { Table, Header, Segment, Label } from 'semantic-ui-react'

export default function ResultsTable2({results}) {
    const rows = results.map(((result, index) => {
        let color;
        if (index === 0) {
            color='blue';
        } else if (index === 1) {
            color='blue';
        } else if (index === 2) {
            color='blue';
        }
        return (
            <Table.Row key={ index }>
                <Table.Cell>
                    <Label ribbon color={color}>{ index + 1 }</Label>
                </Table.Cell>
                <Table.Cell>{ result.cryptocurrency_pair }</Table.Cell>
                <Table.Cell>{ result.max_amount }</Table.Cell>
                <Table.Cell>{ result.raw_yield }</Table.Cell>
                <Table.Cell>{ result.normalized_yield }</Table.Cell>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment>
                <Header>Best-yielding cryptocurrency pairs are: </Header>
                <Table>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Ranking</Table.HeaderCell>
                            <Table.HeaderCell>Currency Pair</Table.HeaderCell>
                            <Table.HeaderCell>Amount</Table.HeaderCell>
                            <Table.HeaderCell>Yield raw</Table.HeaderCell>
                            <Table.HeaderCell>Yield normalized by vol</Table.HeaderCell>
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