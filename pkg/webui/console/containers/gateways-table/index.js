// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import bind from 'autobind-decorator'

import Status from '@ttn-lw/components/status'

import FetchTable from '@ttn-lw/containers/fetch-table'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { checkFromState, mayCreateGateways } from '@console/lib/feature-checks'

import { getGatewaysList } from '@console/store/actions/gateways'

import {
  selectGateways,
  selectGatewaysTotalCount,
  selectGatewaysFetching,
  selectGatewaysError,
} from '@console/store/selectors/gateways'

const headers = [
  {
    name: 'ids.gateway_id',
    displayName: sharedMessages.id,
    width: 24,
    sortable: true,
    sortKey: 'gateway_id',
  },
  {
    name: 'name',
    displayName: sharedMessages.name,
    width: 22,
    sortable: true,
  },
  {
    name: 'ids.eui',
    displayName: sharedMessages.gatewayEUI,
    width: 20,
    sortable: true,
    sortKey: 'gateway_eui',
  },
  {
    name: 'frequency_plan_id',
    displayName: sharedMessages.frequencyPlan,
    width: 16,
  },
  {
    name: 'status',
    width: 18,
    displayName: sharedMessages.status,
    render(status) {
      let indicator = 'unknown'
      let label = sharedMessages.unknown

      if (status === 'connected') {
        indicator = 'good'
        label = sharedMessages.connected
      } else if (status === 'disconnected') {
        indicator = 'bad'
        label = sharedMessages.disconnected
      } else if (status === 'other-cluster') {
        indicator = 'unknown'
        label = sharedMessages.otherCluster
      } else if (status === 'unknown') {
        indicator = 'unknown'
        label = sharedMessages.unknown
      }

      return <Status status={indicator} label={label} />
    },
  },
]

export default class GatewaysTable extends React.Component {
  constructor(props) {
    super(props)

    this.getGatewaysList = params =>
      getGatewaysList(
        params,
        ['name', 'description', 'frequency_plan_id', 'gateway_server_address'],
        { withStatus: true },
      )
  }

  @bind
  baseDataSelector(state) {
    return {
      gateways: selectGateways(state),
      totalCount: selectGatewaysTotalCount(state),
      fetching: selectGatewaysFetching(state),
      error: selectGatewaysError(state),
      mayAdd: checkFromState(mayCreateGateways, state),
    }
  }

  render() {
    return (
      <FetchTable
        entity="gateways"
        addMessage={sharedMessages.addGateway}
        headers={headers}
        getItemsAction={this.getGatewaysList}
        baseDataSelector={this.baseDataSelector}
        tableTitle={<Message content={sharedMessages.gateways} />}
        searchable
        {...this.props}
      />
    )
  }
}
