import React from 'react';
import { Form, Table, Segment, Label } from 'semantic-ui-react'

export default function MyPortfolioTable({results}) {
    const rows = results.map(((result, index) => {
        let color='grey';
        return (
            <Table.Row key={ index }>
                <td><Label class="ui horizontal label" color={color}>{ index + 1 }</Label></td>
                <td>{ result.tokenorpair }</td>
                <td class="right aligned">{ result.amount }</td>
                <td class="right aligned">{ result.percentageofportfolio }</td>
                <td class="right aligned">{ result.risk_setting }</td>
            </Table.Row>
        );
    }));
    return (
        <div className="ui container">
            <Segment class="ui inverted segment">
                <div className="recommended-float">
                    <div className="recommended-header">My Current Portfolio</div>
                    <div className="recommended-reset">
                        <Form className="tableButton">
                            <div class="ui button">Reset</div>
                        </Form>
                    </div>
                </div>
                <div class="ui basic table">
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell><h3 className="headerTitle">Ranking</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Token/Pair</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Amount</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">% Portfolio</h3></Table.HeaderCell>
                            <Table.HeaderCell><h3 className="headerTitle">Risk Setting</h3></Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>
                        { rows }
                    </Table.Body>
                </div>
            </Segment>
        </div>
    );
}
